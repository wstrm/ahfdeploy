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
	Region() (region string)
	Provider() (providerID int)
	Push(image string) (err error)
	NewService(serviceName string, containerName string) (result string, err error)
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

func (a *awsClient) NewService(serviceName string, containerName string) (result string, err error) {
	client := a.client

	input := &ecs.CreateServiceInput{
		// TODO: should not hard code to 1 instance
		DesiredCount:   aws.Int64(1),
		ServiceName:    aws.String(serviceName),
		TaskDefinition: aws.String(containerName),
		// TODO: create cluster
		Cluster: aws.String("test-cluster"),
	}

	awsResult, awsErr := client.CreateService(input)
	if awsErr != nil {
		err = awsErr
		return
	}

	result = awsResult.String()
	return
}

func newAWSClient(region string) (client *awsClient, err error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return
	}

	_, err = sess.Config.Credentials.Get()
	if err != nil {
		return
	}

	client = &awsClient{
		region: region,
		client: ecs.New(sess),
	}

	log.Println(client.client.ListTaskDefinitions(&ecs.ListTaskDefinitionsInput{}))

	return
}

func NewClient(providerID int, region string) (client Client, err error) {
	switch providerID {
	case provider.AWS:
		client, err = newAWSClient(region)
	default:
		err = fmt.Errorf("unknown provider ID: %d", providerID)
	}

	return
}
