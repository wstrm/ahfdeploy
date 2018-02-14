package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/willeponken/d0020e-arrowhead/provider"
)

type Client struct {
	region  string
	cluster string
	client  *ecs.ECS
}

type genericResultSig interface {
	String() string
}

type genericErrorSig interface {
	error
}

func generalizeFuncReturn(result genericResultSig, err genericErrorSig) (string, error) {
	return result.String(), err
}

func (a *Client) getClient() *ecs.ECS {
	return a.client
}

func (a *Client) listClusters() (list *ecs.ListClustersOutput, err error) {
	client := a.getClient()
	list, err = client.ListClusters(&ecs.ListClustersInput{})
	return
}

func (a *Client) clusterExists(name string) (exists bool) {
	list, err := a.listClusters()
	if err != nil {
		return
	}

	for _, cluster := range list.ClusterArns {
		if *cluster == name {
			exists = true
			break
		}
	}

	return
}

func (a *Client) Region() string {
	return a.region
}

func (_ *Client) Provider() int {
	return provider.AWS
}

func (a *Client) Push(image string) error {
	log.Println(image)
	return nil
}

func (a *Client) createCluster(clusterName string) (result string, err error) {
	client := a.getClient()

	input := &ecs.CreateClusterInput{
		ClusterName: aws.String(clusterName),
	}

	result, err = generalizeFuncReturn(client.CreateCluster(input))
	return
}

func (a *Client) Run(serviceName, clusterName, containerName string) (result string, err error) {
	client := a.getClient()

	// create a new cluster if the provided does not exist
	if !a.clusterExists(clusterName) {
		result, err = a.createCluster(clusterName)
		if err != nil {
			return
		}
	}

	input := &ecs.CreateServiceInput{
		// TODO: should not hard code to 1 instance
		DesiredCount:   aws.Int64(1),
		ServiceName:    aws.String(serviceName),
		TaskDefinition: aws.String(containerName),
		Cluster:        aws.String(clusterName),
	}

	result, err = generalizeFuncReturn(client.CreateService(input))
	return
}

func NewClient(region string) (client *Client, err error) {
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

	client = &Client{
		region: region,
		client: ecs.New(sess),
	}

	return
}
