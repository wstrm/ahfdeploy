ahfdeploy
=========
Arrowhead Framework automatic cloud deployment project in the D0020E course.

## Setup development environment
### Prerequisites
 * Download and install Go (https://golang.org/dl/)
 * (optional) Install GoLand IDE (https://www.jetbrains.com/go/)

### Dependency management
 * Install the dependency manager `dep` with `go get -u github.com/golang/dep/cmd/dep`
 * Run `dep ensure` and you're good to go!

## Development guidelines
### Lint and format
Before you commit, the affected files should be run through `goimports`. Bonus: Also run `go vet`.

 * goimports: `go get golang.org/x/tools/cmd/goimports`

### Tests
Write tests for everything.

### Branches
Develop inside your own branch until you're done, then pull request.

### Commits
Commits should be formatted with primary affected package as prefix, a short descriptive one liner and then an optional description the context and what the change does. Also use GitHub's `fixes #123` feature for closing issues.

Example:
```
math: improve Sin, Cos and Tan precision for very large arguments

The existing implementation has poor numerical properties for
large arguments, so use the McGillicutty algorithm to improve
accuracy above 1e10.

The algorithm is described at http://wikipedia.org/wiki/McGillicutty_Algorithm

Fixes #159

# Please enter the commit message for your changes. Lines starting
# with '#' will be ignored, and an empty message aborts the commit.
# On branch foo
# Changes not staged for commit:
#	modified:   editedfile.go
#
```

### Code style
Take a look at: https://github.com/golang/go/wiki/CodeReviewComments

## Project status
| Build status | Test coverage |
|:------------:|:-------------:|
| [![Build Status](https://travis-ci.org/willeponken/d0020e-arrowhead.svg?branch=master)](https://travis-ci.org/willeponken/d0020e-arrowhead) | [![Coverage Status](https://coveralls.io/repos/github/willeponken/d0020e-arrowhead/badge.svg?branch=master)](https://coveralls.io/github/willeponken/d0020e-arrowhead?branch=master) |

## Documentation
### Official Arrowhead Framework documentation
Link to official AHF repo: https://forge.soa4d.org/scm/?group_id=58
