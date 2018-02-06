package session

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/willeponken/d0020e-arrowhead/provider"
)

type Client interface {
	Region() string
	Provider() int
	Push(image string) error
}

type awsClient struct {
	region string
	client *ecs.ECS
}

func (a *awsClient) Region() string {
	return a.region
}

func (a *awsClient) Provider() int {
	return provider.AWS
}

func (a *awsClient) Push(image string) error {
	log.Println(image)
	return nil
}

func newAWSClient(region string) (client *awsClient, err error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return
	}

	client = &awsClient{
		region: region,
		client: ecs.New(sess),
	}

	return
}

func New(providerID int, region string) (client *Client, err error) {
	switch providerID {
	case provider.AWS:
		client, err = newAWSClient(provider)
	default:
		err = fmt.Errorf("unknown provider ID: %d", providerID)
	}

	return
}
