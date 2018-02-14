package aws

import (
	"bytes"
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	types "github.com/moby/moby/api/types"
	docker "github.com/moby/moby/client"
	"github.com/willeponken/d0020e-arrowhead/provider"
)

type Client struct {
	region       string
	cluster      string
	ecsClient    *ecs.ECS
	dockerClient *docker.Client
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

func (a *Client) getECSClient() *ecs.ECS {
	return a.ecsClient
}

func (a *Client) getDockerClient() *docker.Client {
	return a.dockerClient
}

func (a *Client) listClusters() (list *ecs.ListClustersOutput, err error) {
	client := a.getECSClient()
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

func (a *Client) createCluster(clusterName string) (result string, err error) {
	client := a.getECSClient()

	input := &ecs.CreateClusterInput{
		ClusterName: aws.String(clusterName),
	}

	result, err = generalizeFuncReturn(client.CreateCluster(input))
	return
}

func (a *Client) putImage(imageURI, registryAuth string) error {
	client := a.getDockerClient()

	result, err := client.ImagePush(context.Background(), imageURI, &types.ImagePushOptions{
		RegistryAuth: registryAuth, // RegistryAuth is the base64 encoded credentials for the registry
	})
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(result)
	log.Println(buf.String())

	return nil
}

func (a *Client) Region() string {
	return a.region
}

func (_ *Client) Provider() int {
	return provider.AWS
}

func (a *Client) Upload(image string) error {
	log.Println(image)
	return nil
}

func (a *Client) Upload(image string) error {
	return a.putImage(image, "NOTAREALYAUTHKEYSHOULDBERETRIEVEDSOMEWHEREIDUNNO")
}

func (a *Client) Run(serviceName, clusterName, containerName string) (result string, err error) {
	client := a.getECSClient()

	// create a new cluster if the provided does not exist
	if !a.clusterExists(clusterName) {
		result, err = a.createCluster(clusterName)
		if err != nil {
			return
		}
	}

	input := &ecs.CreateServiceInput{
		// TODO: should not hardcode to 1 instance
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

	dockerClient, err := docker.NewEnvClient()
	if err != nil {
		return
	}

	client = &Client{
		region:       region,
		ecsClient:    ecs.New(sess),
		dockerClient: dockerClient,
	}

	return
}
