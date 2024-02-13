package main

// TargetGroup represents a single group of targets with labels.
type TargetGroup struct {
	Targets []string          `yaml:"targets"`
	Labels  map[string]string `yaml:"labels"`
}

// PrometheusConfig is a slice of TargetGroup, allowing multiple groups of targets.
type PrometheusConfig []TargetGroup
