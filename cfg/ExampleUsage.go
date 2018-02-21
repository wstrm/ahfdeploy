package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Title string
	Base base
	AwsConf awsConf
	GoogleConf googleConf
}

type base struct {
	Provider string
	DateChanged string
}

type awsConf struct {
	UserID int
	Token string
	Region string
}

type googleConf struct {
	UserID int
	Token string
	Region string
}

// Use as normal struct, config.X
func main() {
	var config tomlConfig
	if _, err := toml.DecodeFile("AHFDeployConf.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Title: %s\n", config.Title)
	fmt.Printf("Provider: %s\n", config.Base.Provider)
	fmt.Printf("AWS: %d, %s, %s\n", config.AwsConf.UserID, 
		config.AwsConf.Token, config.AwsConf.Region)
	fmt.Printf("Google: %d, %s, %s\n", config.GoogleConf.UserID,
		config.GoogleConf.Token, config.GoogleConf.Region)
	fmt.Printf("Last changed: %s\n", config.Base.DateChanged)
}