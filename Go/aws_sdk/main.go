package main

import (
	"context"
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("failed to load AWS configuration: %v", err)
	}

	// Setup ec2 client
	ec2Client := ec2.NewFromConfig(cfg)

	// Take user input in menu style for SDK actions
	for {
		printMainMenu()
		action := collectUserInput()
		switch action {
		case "1":
			err = ListEc2Instances(ec2Client, cfg, ctx)
			if err != nil {
				log.Fatalf("failed to list EC2 instances: %v", err)
			}
		case "Q":
			return
		default:
			fmt.Println("Invalid action. Please try again.")
		}
		fmt.Println()
	}
}

func printMainMenu() {
	menu := `1. List EC2 Instances
Q. Exit
`
	fmt.Printf(menu)
}

func collectUserInput() string {
	validActions := []string{"1", "Q"}
	var action string
	for {
		fmt.Printf("choice:~$ ")
		_, err := fmt.Scan(&action)
		action = strings.ToUpper(action)
		if err != nil {
			log.Fatalf("failed to read user input: %v", err)
		}
		if slices.Contains(validActions, action) {
			break
		}
		fmt.Println("Invalid action. Please try again.")
	}
	return action
}
