package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/willeponken/ahfdeploy/deploy"
	providers "github.com/willeponken/ahfdeploy/provider"
)

type awsConfig struct {
	Repository string
	Cluster    string
	Service    string
	Task       string
	Region     string
}

type baseConfig struct {
	Provider string
	AWS      awsConfig `toml:"aws"`
}

type baseFlags struct {
	baseConfig
}

func main() {
	var config baseConfig
	var configFilepath string

	flag.StringVar(&configFilepath, "config", "ahfdeploy.toml", "Configuration filepath")
	flag.StringVar(&config.Provider, "provider", "", "Service provider")
	flag.StringVar(&config.AWS.Region, "aws-region", "", "Region for AWS")
	flag.StringVar(&config.AWS.Repository, "aws-repository", "", "Repository for container upload to AWS ECR")
	flag.StringVar(&config.AWS.Cluster, "aws-cluster", "ahfdeploy-default-cluster", "Cluster name for AWS ECS")
	flag.StringVar(&config.AWS.Service, "aws-service", "ahfdeploy-default-service", "Service name for AWS ECS")
	flag.StringVar(&config.AWS.Task, "aws-task", "ahfdeploy-default-task", "Task Definition name for AWS ECS")
	flag.Parse()

	if _, err := toml.DecodeFile(configFilepath, &config); err != nil {
		log.Fatalln(err)
	}

	if config.Provider == "" {
		log.Fatalln("Please define a provider.")
	}

	switch config.Provider {
	case "aws":
		if config.AWS.Repository == "" {
			log.Fatalln("Please define a AWS repository.")
		}

		client, err := deploy.NewClient(providers.AWS, config.AWS.Region)
		if err != nil {
			log.Fatalln(err)
		}

		resultUpload, err := client.Upload(config.AWS.Repository)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(resultUpload)

		resultRun, err := client.Create(config.AWS.Service, config.AWS.Cluster, config.AWS.Task)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(resultRun)
	default:
		log.Fatalf("Unknown provider: %s", config.Provider)
	}
}
