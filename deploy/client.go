package deploy

import (
	"fmt"

	"github.com/willeponken/ahfdeploy/deploy/aws"
	"github.com/willeponken/ahfdeploy/provider"
)

type Client interface {
	Region() (region string)
	Provider() (providerID int)
	Upload(imageURI string) (result string, err error)
	Run(serviceName, clusterName, containerName string) (result string, err error)
}

func NewClient(providerID int, region string) (client Client, err error) {
	switch providerID {
	case provider.AWS:
		client, err = aws.NewClient(region)
	default:
		err = fmt.Errorf("unknown provider ID: %d", providerID)
	}

	return
}
