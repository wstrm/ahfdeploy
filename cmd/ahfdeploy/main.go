package main

import (
	"log"

	"github.com/willeponken/d0020e-arrowhead/deploy"
	"github.com/willeponken/d0020e-arrowhead/provider"
)

func main() {
	client, err := deploy.NewClient(provider.AWS, "us-west-2")
	if err != nil {
		log.Fatalln(err)
	}

	result, err := client.Run("ecs-test-service", "test-cluster", "hello_world")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(result)
}
