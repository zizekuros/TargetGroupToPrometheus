package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func getPublicIPsFromTargetGroup(targetGroupARN string, awsProfile string, accessKeyID string, secretAccessKey string) ([]string, error) {
	var cfg aws.Config
	var err error

	if accessKeyID != "" && secretAccessKey != "" {
		// Use static credentials if both access key and secret key are provided
		creds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, ""))
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds))
	} else if awsProfile != "" {
		// Use the specified AWS profile if provided
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(awsProfile))
	} else {
		// Load the default configuration
		cfg, err = config.LoadDefaultConfig(context.TODO())
	}

	if err != nil {
		return nil, fmt.Errorf("error loading AWS configuration: %w", err)
	}

	// Create an ELB client
	elbv2Client := elasticloadbalancingv2.NewFromConfig(cfg)

	// Describe target health to get instance IDs
	targetHealthOutput, err := elbv2Client.DescribeTargetHealth(context.TODO(), &elasticloadbalancingv2.DescribeTargetHealthInput{
		TargetGroupArn: aws.String(targetGroupARN),
	})
	if err != nil {
		return nil, fmt.Errorf("error describing target health: %w", err)
	}

	// Extract instance IDs from the targets
	instanceIDs := make([]string, len(targetHealthOutput.TargetHealthDescriptions))
	for i, thd := range targetHealthOutput.TargetHealthDescriptions {
		instanceIDs[i] = *thd.Target.Id
	}

	// Create an EC2 client
	ec2Client := ec2.NewFromConfig(cfg)

	// Describe instances to get public IPs
	publicIPs := []string{}
	if len(instanceIDs) > 0 {
		describeInstancesOutput, err := ec2Client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
			InstanceIds: instanceIDs,
		})
		if err != nil {
			return nil, fmt.Errorf("error describing EC2 instances: %w", err)
		}

		for _, reservation := range describeInstancesOutput.Reservations {
			for _, instance := range reservation.Instances {
				if instance.PublicIpAddress != nil {
					publicIPs = append(publicIPs, *instance.PublicIpAddress)
				}
			}
		}
	}

	return publicIPs, nil
}
