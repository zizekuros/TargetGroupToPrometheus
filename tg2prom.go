package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Command-line flags
	targetGroupARN := flag.String("target-group-arn", "", "The ARN of the target group (required)")
	awsProfile := flag.String("aws-profile", "", "AWS profile to use for credentials and configuration (default is used if not provided)")
	awsAccessKeyID := flag.String("aws-access-key-id", "", "AWS Access Key ID")
	awsSecretAccessKey := flag.String("aws-secret-access-key", "", "AWS Secret Access Key")
	listeningPort := flag.Int("listening-port", 0, "The listening port for scraping the metrics (required)")
	jobName := flag.String("job-name", "", "The desired name of the job label (required)")
	outputFileName := flag.String("output", "", "The desired output file name (required)")

	// Custom usage message for the --help flag
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *targetGroupARN == "" || *listeningPort == 0 || *jobName == "" || *outputFileName == "" {
		fmt.Println("Error: All parameters are required.")
		flag.Usage()
		os.Exit(1)
	}

	list, err := getPublicIPsFromTargetGroup(*targetGroupARN, *awsProfile, *awsAccessKeyID, *awsSecretAccessKey)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	err = generatePrometheusConfig(list, *jobName, *listeningPort, *outputFileName)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Successfully generated Prometheus targets configuration at '%s'.\n", *outputFileName)

}
