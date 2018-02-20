package aws

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecs"
	types "github.com/docker/docker/api/types"
	docker "github.com/docker/docker/client"
	"github.com/willeponken/d0020e-arrowhead/provider"
)

type Client struct {
	region       string
	cluster      string
	ecsClient    *ecs.ECS
	ecrClient    *ecr.ECR
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

func (a *Client) getECRClient() *ecr.ECR {
	return a.ecrClient
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

func decodeECRToken(token string) (user, password string, err error) {
	authPairBytes, err := base64.RawStdEncoding.DecodeString(token)
	if err != nil {
		return
	}

	authPairBytesSpl := bytes.Split(authPairBytes, []byte(":"))
	if len(authPairBytesSpl) > 2 {
		panic("auth pair contains multiple colons")
	}

	if len(authPairBytesSpl) != 2 {
		panic("invalid auth pair")
	}

	user = fmt.Sprintf("%s", authPairBytesSpl[0])
	password = fmt.Sprintf("%s", authPairBytesSpl[1])

	return
}

func (a *Client) getRegistryAuthorization() (registryAuth string, err error) {
	client := a.getECRClient()

	result, err := client.GetAuthorizationToken(&ecr.GetAuthorizationTokenInput{})

	if err != nil {
		return
	}

	var token string
	for _, auth := range result.AuthorizationData {
		token = *auth.AuthorizationToken

		if token != "" {
			break
		}
	}

	if token == "" {
		err = errors.New("no authorization token found in ECR response")
		return
	}

	user, password, err := decodeECRToken(token)
	if err != nil {
		return
	}

	registryAuthBytes, err := json.Marshal(struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}{
		Username: user,
		Password: password,
		Email:    "none",
	})
	if err != nil {
		return
	}

	registryAuth = base64.StdEncoding.EncodeToString(registryAuthBytes)
	return
}

func (a *Client) Region() string {
	return a.region
}

func (_ *Client) Provider() int {
	return provider.AWS
}

func (a *Client) Upload(imageURI string) (string, error) {
	client := a.getDockerClient()

	registryAuth, err := a.getRegistryAuthorization()
	if err != nil {
		return "", err
	}

	result, err := client.ImagePush(context.Background(), imageURI, types.ImagePushOptions{
		RegistryAuth: registryAuth, // RegistryAuth is the base64 encoded credentials for the registry
	})
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(result)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
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
		ecrClient:    ecr.New(sess),
		dockerClient: dockerClient,
	}

	return
}
