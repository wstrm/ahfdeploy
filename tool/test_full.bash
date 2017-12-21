#!/usr/bin/env bash

set -e

pkgName=$1 # repo-name
pkgImport=$2 # example.com/user/repo-name/cmd/daemon

echo "Testing repo: $pkgName ($pkgImport)"

dir=$(dirname $0)

local_dependencies() {
	$dir/godeps.bash $pkgImport $pkgName $1
}

echo "# Golint"
golint -min_confidence 0.0 $(local_dependencies /...)

echo "# Go vet"
go vet $(local_dependencies /...)

echo "# Go test"
# Quiet cgo when compiling therecipe/qt
export CGO_CPPFLAGS="-Wno-unused-variable -Wno-unused-parameter -Wno-return-type"
echo 'mode: atomic' > coverage.txt && local_dependencies | xargs -n1 -I{} sh -c 'echo "> {}"; go test -tags test -race -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp