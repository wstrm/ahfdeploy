package main

import (
	"flag"
	"log"
	"os"

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
	Token      string
	Secret     string
	ID         string `toml:"id"`
}

type baseConfig struct {
	Provider string
	AWS      awsConfig `toml:"aws"`
}

type baseFlags struct {
	baseConfig
}

func exitWithHelp(code int, message string) {
	log.Println(message)
	flag.Usage()
	os.Exit(code)
}

func main() {
	var config baseConfig
	var configFilepath string

	flag.StringVar(&configFilepath, "config", "ahfdeploy.toml", "Configuration filepath")
	flag.StringVar(&config.Provider, "provider", "", "Service provider")
	flag.StringVar(&config.AWS.Token, "aws-token", "", "Token for AWS")
	flag.StringVar(&config.AWS.Secret, "aws-secret", "", "Secret for AWS")
	flag.StringVar(&config.AWS.ID, "aws-id", "", "ID for AWS")
	flag.StringVar(&config.AWS.Region, "aws-region", "", "Region for AWS")
	flag.StringVar(&config.AWS.Repository, "aws-repository", "", "Repository for container upload to AWS ECR")
	flag.StringVar(&config.AWS.Cluster, "aws-cluster", "ahfdeploy-default-cluster", "Cluster name for AWS ECS")
	flag.StringVar(&config.AWS.Service, "aws-service", "ahfdeploy-default-service", "Service name for AWS ECS")
	flag.StringVar(&config.AWS.Task, "aws-task", "ahfdeploy-default-task", "Task Definition name for AWS ECS")
	flag.Parse()

	// Load configuration file, if it doesn't exist or is invalid, we'll only use flags.
	toml.DecodeFile(configFilepath, &config)

	if config.Provider == "" {
		exitWithHelp(1, "Please define a provider.")
	}

	switch config.Provider {
	case "aws":
		if config.AWS.Repository == "" {
			exitWithHelp(1, "Please define a AWS repository.")
		}

		var (
			client deploy.Client
			err    error
		)
		if config.AWS.Token == "" || config.AWS.Secret == "" || config.AWS.ID == "" {
			client, err = deploy.NewClient(providers.AWS, config.AWS.Region)
		} else {
			client, err = deploy.NewClientWithCredentials(providers.AWS,
				config.AWS.Region,
				config.AWS.ID,
				config.AWS.Secret,
				config.AWS.Token)
		}
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
