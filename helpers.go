package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func generatePrometheusConfig(publicIPs []string, jobName string, listeningPort int, outputFileName string) error {
	// Prepare the targets with the port
	var targets []string
	for _, ip := range publicIPs {
		target := fmt.Sprintf("%s:%d", ip, listeningPort)
		targets = append(targets, target)
	}

	// Create the Prometheus configuration structure
	config := PrometheusConfig{
		Targets: targets,
	}
	config.Labels.Job = jobName

	// Marshal the configuration to YAML
	data, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("error marshaling YAML: %w", err)
	}

	// Write the YAML to the specified output file
	err = os.WriteFile(outputFileName, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file '%s': %w", outputFileName, err)
	}

	return nil
}
