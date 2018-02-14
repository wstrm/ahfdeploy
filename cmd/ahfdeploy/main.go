package main

import (
	"log"

	"github.com/willeponken/d0020e-arrowhead/provider"
	"github.com/willeponken/d0020e-arrowhead/session"
)

func main() {
	client, err := session.NewClient(provider.AWS, "us-west-2")
	if err != nil {
		log.Fatalln(err)
	}

	result, err := client.NewService("ecs-test-service", "hello_world")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(result)
}
