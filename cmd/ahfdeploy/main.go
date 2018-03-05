package main

import (
	"log"
	"flag"
  
	"github.com/willeponken/ahfdeploy/deploy"
	providers "github.com/willeponken/ahfdeploy/provider"
)

func main() {
	var (
		region string
		provider string
	)

	flag.StringVar(&region, "region", "", "Unique region for provider")
	flag.StringVar(&provider, "provider", "", "Service provider")
	flag.Parse()

	if region == "" {
		log.Fatalln("Please define a region.")
	}

	if provider == "" {
		log.Fatalln("Please define a provider.")
	}

	var p int
	switch provider {
	case "aws":
		p = providers.AWS
	default:
		log.Fatalf("Unknown provider: %s", provider)
	}

	client, err := deploy.NewClient(p, region)
	if err != nil {
		log.Fatalln(err)
	}

	resultUpload, err := client.Upload("059336174526.dkr.ecr.us-west-2.amazonaws.com/willeponken:latest")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resultUpload)

	resultRun, err := client.Create("ecs-test-service", "test-cluster", "hello_world")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resultRun)
}
