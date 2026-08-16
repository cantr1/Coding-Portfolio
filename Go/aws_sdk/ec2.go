package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type InstanceSummary struct {
	ID               string
	Name             string
	Type             string
	State            string
	AvailabilityZone string
	PublicIP         string
	VpcID            string
}

func stringValue(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}

func tagValue(tags []types.Tag, key string) string {
	for _, tag := range tags {
		if stringValue(tag.Key) == key {
			return stringValue(tag.Value)
		}
	}

	return ""
}

func StartEc2Instance(client *ec2.Client, ctx context.Context, instanceID string) error {
	_, err := client.StartInstances(ctx, &ec2.StartInstancesInput{
		InstanceIds: []string{instanceID},
	})
	if err != nil {
		return err
	}

	return nil
}

func StopEc2Instance(client *ec2.Client, ctx context.Context, instanceID string) error {
	_, err := client.StopInstances(ctx, &ec2.StopInstancesInput{
		InstanceIds: []string{instanceID},
	})
	if err != nil {
		return err
	}

	return nil
}

func ListEc2Instances(client *ec2.Client, ctx context.Context) error {
	paginator := ec2.NewDescribeInstancesPaginator(client, &ec2.DescribeInstancesInput{})

	var instances []InstanceSummary

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		for _, reservation := range page.Reservations {
			for _, instance := range reservation.Instances {
				summary := InstanceSummary{
					ID:               stringValue(instance.InstanceId),
					Type:             string(instance.InstanceType),
					Name:             tagValue(instance.Tags, "Name"),
					State:            string(instance.State.Name),
					AvailabilityZone: stringValue(instance.Placement.AvailabilityZone),
					PublicIP:         stringValue(instance.PublicIpAddress),
					VpcID:            stringValue(instance.VpcId),
				}

				instances = append(instances, summary)
			}
		}
	}

	printEc2Instances(instances)

	return nil
}

func printEc2Instances(instances []InstanceSummary) {
	fmt.Printf(
		"%-22s %-24s %-14s %-12s %-16s %-16s %-22s\n",
		"ID",
		"Name",
		"Type",
		"State",
		"AZ",
		"Public IP",
		"VPC ID",
	)

	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------")

	for _, instance := range instances {
		fmt.Printf(
			"%-22s %-24s %-14s %-12s %-16s %-16s %-22s\n",
			instance.ID,
			instance.Name,
			instance.Type,
			instance.State,
			instance.AvailabilityZone,
			instance.PublicIP,
			instance.VpcID,
		)
	}
}
