package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	instanceID := os.Args[1]

	listInstances(instanceID)
}

func listInstances(instanceID string) {
	awsConnect, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-north-1"),
		Credentials: credentials.NewSharedCredentials("", "terraform"),
	},
	)
	if err != nil {
		fmt.Println(err)
	}

	ec2sess := ec2.New(awsConnect)

	instanceInfo := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	}
	fmt.Println("Listing EC2 Instance info...")

	fmt.Println(ec2sess.DescribeInstances(instanceInfo))
}
