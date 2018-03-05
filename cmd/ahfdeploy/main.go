package main

import (
	"log"

	"github.com/willeponken/ahfdeploy/deploy"
	"github.com/willeponken/ahfdeploy/provider"
)

func main() {
	client, err := deploy.NewClient(provider.AWS, "us-west-2")
	if err != nil {
		log.Fatalln(err)
	}

	resultUpload, err := client.Upload("059336174526.dkr.ecr.us-west-2.amazonaws.com/willeponken:latest")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resultUpload)

	resultRun, err := client.Run("ecs-test-service", "test-cluster", "hello_world")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resultRun)
}
