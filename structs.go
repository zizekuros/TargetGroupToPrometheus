package main

// PrometheusConfig represents the structure of the Prometheus targets YAML file.
type PrometheusConfig struct {
	Targets []string `yaml:"targets"`
	Labels  struct {
		Job string `yaml:"job"`
	} `yaml:"labels"`
}
