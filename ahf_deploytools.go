package main

import (
	dockerclient "github.com/docker/docker/client"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/docker/docker/api/types"
	"github.com/aws/aws-sdk-go/aws"
	"context"
	"fmt"
)

func main(){
	//createRepo("ahf-test2");
	//createCluster("testcluster");
	//deleteRepo("ahf-test");
	//deleteCluster("testcluster");
	//runTask("default", "console-sample-app-static:2")
	//stopTask("default","5634a3db-f7ff-4a49-b27d-8764307df319")
	//getToken();
	//doPush("554348295977.dkr.ecr.eu-west-2.amazonaws.com/ahf-test");
}

/*
	Stops a running task
 */
func stopTask(cluster string, task string){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
	})
	svc := ecs.New(sess)
	input := &ecs.StopTaskInput{
		Cluster:        aws.String(cluster),
		Task: aws.String(task),
	}

	result, err := svc.StopTask(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException:
				fmt.Println(ecs.ErrCodeServerException, aerr.Error())
			case ecs.ErrCodeClientException:
				fmt.Println(ecs.ErrCodeClientException, aerr.Error())
			case ecs.ErrCodeInvalidParameterException:
				fmt.Println(ecs.ErrCodeInvalidParameterException, aerr.Error())
			case ecs.ErrCodeClusterNotFoundException:
				fmt.Println(ecs.ErrCodeClusterNotFoundException, aerr.Error())
			case ecs.ErrCodeUnsupportedFeatureException:
				fmt.Println(ecs.ErrCodeUnsupportedFeatureException, aerr.Error())
			case ecs.ErrCodePlatformUnknownException:
				fmt.Println(ecs.ErrCodePlatformUnknownException, aerr.Error())
			case ecs.ErrCodePlatformTaskDefinitionIncompatibilityException:
				fmt.Println(ecs.ErrCodePlatformTaskDefinitionIncompatibilityException, aerr.Error())
			case ecs.ErrCodeAccessDeniedException:
				fmt.Println(ecs.ErrCodeAccessDeniedException, aerr.Error())
			case ecs.ErrCodeBlockedException:
				fmt.Println(ecs.ErrCodeBlockedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

/*
	Runs an active, stopped task
 */
func runTask(cluster string, task string){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
	})
	svc := ecs.New(sess)
	input := &ecs.RunTaskInput{
		Cluster:        aws.String(cluster),
		TaskDefinition: aws.String(task),
	}

	result, err := svc.RunTask(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException:
				fmt.Println(ecs.ErrCodeServerException, aerr.Error())
			case ecs.ErrCodeClientException:
				fmt.Println(ecs.ErrCodeClientException, aerr.Error())
			case ecs.ErrCodeInvalidParameterException:
				fmt.Println(ecs.ErrCodeInvalidParameterException, aerr.Error())
			case ecs.ErrCodeClusterNotFoundException:
				fmt.Println(ecs.ErrCodeClusterNotFoundException, aerr.Error())
			case ecs.ErrCodeUnsupportedFeatureException:
				fmt.Println(ecs.ErrCodeUnsupportedFeatureException, aerr.Error())
			case ecs.ErrCodePlatformUnknownException:
				fmt.Println(ecs.ErrCodePlatformUnknownException, aerr.Error())
			case ecs.ErrCodePlatformTaskDefinitionIncompatibilityException:
				fmt.Println(ecs.ErrCodePlatformTaskDefinitionIncompatibilityException, aerr.Error())
			case ecs.ErrCodeAccessDeniedException:
				fmt.Println(ecs.ErrCodeAccessDeniedException, aerr.Error())
			case ecs.ErrCodeBlockedException:
				fmt.Println(ecs.ErrCodeBlockedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

/*
	Retrieves a 12-hour auth token from aws ecr to be used with docker login
 */
func getToken(){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
	})
	svc := ecr.New(sess)
	input := &ecr.GetAuthorizationTokenInput{}

	result, err := svc.GetAuthorizationToken(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

/*
	A docker push function, currently not implemented(docker needs to be signed in with token first)
 */
func doPush(target string) {
	envCli, err := dockerclient.NewEnvClient()
	if err != nil {
		panic(err)
	}

	rc, err := envCli.ImagePush(
		context.Background(),
		target,
		types.ImagePushOptions{All: true,
			RegistryAuth:"QVdTOmV5SndZWGxzYjJGa0lqb2lURlZaWjNadksxSk1RWEphYVdkMlFpdDRSa2N6YnpsNVdsUlhXbU5xU2xaMWEwVnhibXRRV1ZaWlpuSTNNVU0xVVdSS2NXa3hSRTlXTlZob1lVOUNWa1Z2VVhCNFowNU9LelZGYWt4WVJYazBabkpSVnpkUFRUWmpNVWhaY0hWNkszaHFNV1V3YjJKelQzWklSVEo1VTNoQ1pDdDBUamd2ZEhWbVFUUlhMMjFEYlcxVGVEbE5UMDFxTUhsWk9FTk9jMjkxU1hWYU1reDNTQ3ROYlV0V1RrZFZiMjUxVTBOSU0wMUdibFIzY3pkNUsxUjRZbTQzY0V4NWJHdGljM1JsZVVGWFN6TkJWMVZGZFcxelpUUkpUa1JMZFc0NVJFNHZZa3cwYW1zeGVHWm1hemg1UkRaMmNtUmllbGxhY1hJNFNrMXJOM3BIVjI1VmRIUnpabFpEV2tsaE1YWlRTMVZVVFZFeU5pOTFUelExYkcxVVdsZFlPV2x0WlV0SmFHSk1ZMWgzWjA5amVHRkhZMGhxZDNwRUwzSkhSWFpOVURGT1l6VmxhV3RtYjBkYWRsb3lWVlZ5VFU5WlRrUmhjMk51U2xOb05FOWpOalpGYkdSak5tTnhTR0UyZEZVcmVWTXJOMUZIV25aVVZ6bEVjVXRXU0ZsNlJWSkNkVXQ0WVZwRlQwUlBhMWc0UVdaRlRFcFlSWGczWkZvNE9YVXdVWGhxWVVScE1EbDZXa3h0VFdZMk5rdzVWRWRpVVhrdmQxcHlMek5oWjFBek1HVlVUelJEYkVZM2R6SmpTekpQVFZkd1ZFNDJaV05CWXpWSVdTOVZNRlJNV1Vndk9XVnhVR3REVEhjNE1rdFZOakp3UW1OTFNqaDJXVXgyTWxBd05XeDJlSGxUTmxaWU1YZ3hjMnhtVUZsMlEyVkxkM1V6Y0RkMFZITTBjRVZ2Vm1Gck9VOXZVbk5rVld3MGFVWmxWMm80Y1VaeU9EVlJSVUUxVkdvNGJHeHJlazh3TkhwelIwMW1VWEpoZHk5TWN5OXVWRWR2YkZWSGFubFRlV3c0TUdOQ1JrcHphVFJzT1dsclJFNVlaM0ZqUVZOa1ZtOHJSWFJzWnpOWE1FVmtNRXQzVmsxcGRtSmlVRlpOUjJreFQzUmhibk54VkRWMWMwRTRZblF4TWs5d2IzQlBibWh1WXpsa1drTnZVelZPYVdrMk5uWjFObVJrY2pGVlUweG1iMGRYTWtOelZuWktOWGRIVUZaQlJXdHdiRUpGY1VZNUsyMVhSamswWkZRMldURlBXbFkzWlRGMFREZDRaRzA1U21wSFowTjJSRmd4YkhwMGFrMVNZazF3UVdwVE9HUjRXa1ZTWTBSM05WaHZhbXBrZHpCVVYwaGpVVk4yYVdGS1VVWlFZM1k0UVdSWlNrYzFObmwyYlRsR2FqZEdRbFowYW5kVlltSndUMU56ZVRWeGJVRmtObVI1U21Oc1NIUm5TVEp4WnpWYVdrRXhRa2RCWkVKeFRscG5ZVlEzZGxOdlJHZEdZVzgzWlZNdmJIWTVUbGgxVkRGc1JFSlNLMVpIVUdZdlZsaEhNREZyT1RObVdrNXBSakZTVlc5eVJWUjNRV2R5VFhwWlowWklWVnBKV0RoTFJVWkVUVVZUY0dseWVFNW1NemhyV20xSk0yb3lOMEZxYzJKd1ZrVjVOalJ4YWxOMGRWUm1kRmR0V1VSME0xUlVZMHRtV2tKRldISnZVRVpIV1dwSFpYWlNhV0ZMWW5jMVVUZGpXVkJYYkRGRGIyc3lXVmQzYjBacFV6QlhablZ1VkVNNU9VaElXR0ZyY0d0QllYUTNhVEJMTDB0d1lWcFlTa2hzV0ZkRFNFbDJWVTFUTXpGTFlrUkNORTlFVUV4elZ6aGtlR1oxVXpCdFRTc3pkamswVnpCS1NWaDJVazlVWmswNWJUQlNhbEoxTmxCNUwzQnRhazVRWkZVaUxDSmtZWFJoYTJWNUlqb2lRVkZGUWtGSWFFNUVlbGRXYmxsSGVXeExNM1JWWVd4SVFVNXdiWEp1VVhwaE5YYzBWelJqU0ZsTlZUZ3lkR0pPU1hkQlFVRklOSGRtUVZsS1MyOWFTV2gyWTA1QlVXTkhiMGM0ZDJKUlNVSkJSRUp2UW1kcmNXaHJhVWM1ZHpCQ1FuZEZkMGhuV1VwWlNWcEpRVmRWUkVKQlJYVk5Ra1ZGUkV4TVQwRkZZbnBZSzJOYVoySnFZWGhSU1VKRlNVRTNORGR4TlVKaFFVbEJkamhpWVVkcE9YWlFaRzloVVZsaFRYQjRWMVZwZEVoa04yZHRhV0ZhV1ZCT2FXTjJURU53TWpGc1QxTmxiVnBsY0hsWVpsSjZaRU5SU2pKV2F6STRMM1V3VDB4VmF6MGlMQ0oyWlhKemFXOXVJam9pTWlJc0luUjVjR1VpT2lKRVFWUkJYMHRGV1NJc0ltVjRjR2x5WVhScGIyNGlPakUxTVRZM01EazFNRFY5"})
	if err != nil {
		panic(err)
	}
	fmt.Println(rc)
	defer rc.Close()
}

/*
	Creates a new cluster on aws ecs
 */
func createCluster(cluster string){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
	})
	svc := ecs.New(sess)
	input := &ecs.CreateClusterInput{
		ClusterName: aws.String(cluster),
	}

	result, err := svc.CreateCluster(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException:
				fmt.Println(ecs.ErrCodeServerException, aerr.Error())
			case ecs.ErrCodeClientException:
				fmt.Println(ecs.ErrCodeClientException, aerr.Error())
			case ecs.ErrCodeInvalidParameterException:
				fmt.Println(ecs.ErrCodeInvalidParameterException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

}

/*
	Deletes a cluster on aws ecs
 */
func deleteCluster(cluster string){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
	})
	svc := ecs.New(sess)
	input := &ecs.DeleteClusterInput{
		Cluster: aws.String(cluster),
	}

	result, err := svc.DeleteCluster(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException:
				fmt.Println(ecs.ErrCodeServerException, aerr.Error())
			case ecs.ErrCodeClientException:
				fmt.Println(ecs.ErrCodeClientException, aerr.Error())
			case ecs.ErrCodeInvalidParameterException:
				fmt.Println(ecs.ErrCodeInvalidParameterException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

}

/*
	Deletes a repository on aws ecr
 */
func deleteRepo(repo string){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
	})
	svc := ecr.New(sess)
	input := &ecr.DeleteRepositoryInput{
		Force:          aws.Bool(true),
		RepositoryName: aws.String(repo),
	}

	result, err := svc.DeleteRepository(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeRepositoryNotFoundException:
				fmt.Println(ecr.ErrCodeRepositoryNotFoundException, aerr.Error())
			case ecr.ErrCodeRepositoryNotEmptyException:
				fmt.Println(ecr.ErrCodeRepositoryNotEmptyException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

}

/*
	Creates a new repository on aws ecr
 */
func createRepo(repo string){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
	})
	svc := ecr.New(sess)
	input := &ecr.CreateRepositoryInput{
		RepositoryName: aws.String(repo),
	}

	result, err := svc.CreateRepository(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeRepositoryAlreadyExistsException:
				fmt.Println(ecr.ErrCodeRepositoryAlreadyExistsException, aerr.Error())
			case ecr.ErrCodeLimitExceededException:
				fmt.Println(ecr.ErrCodeLimitExceededException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

}