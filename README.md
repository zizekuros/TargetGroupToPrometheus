# TargetGroupToPrometheus

`TargetGroupToPrometheus` is a CLI tool developed in Go, designed to generate Prometheus target configurations from AWS Elastic Load Balancer (ELB) target groups. 

It retrieves public IP addresses of instances registered with a specified target group and generates a YAML configuration file for Prometheus, facilitating dynamic service discovery.

## Pre-requisites

- Go 1.15 or newer.
- An AWS account and credentials with permissions to access Elastic Load Balancing (ELB) and EC2 services.

## Building the application

1. Clone the repository

```
git clone https://github.com/zizekuros/TargetGroupToPrometheus.git
```

2. Navigate to the project directory:

```sh
cd TargetGroupToPrometheus
```

3. Build the application:

```sh
go build -o tg2prom
```

This command compiles the application and generates an executable named `tg2prom`.

## Running the Application

Run the application by providing the necessary flags. You can view all available options by using `-help`:

```sh
./tg2prom -help
```

To run the tool with AWS credentials and target group details, use:

```sh
./tg2prom -target-group-arn "<target_group_arn>" -listening-port "<port>" -job-name "your_target_name" -output "output_file_name.yaml" [-aws-profile "aws_profile"] [-aws-access-key-id "your_access_key_id" -aws-secret-access-key "your_secret_access_key"]
```

### Flags:
- `target-group-arn`: ARN of the target group (required).
- `listening-port`: The port on which the target instances are listening for Prometheus to scrape metrics (required).
- `job-name`: Name of the job, used as a label in the output file (required).
- `output`: Name of the output YAML file (required).
- `aws-profile`: AWS profile to use for credentials and configuration (optional).
- `aws-access-key-id`: AWS Access Key ID (optional, use if not using `aws-profile`).
- `aws-secret-access-key`: AWS Secret Access Key (optional, use if not using `aws-profile`).

## Example of Config Output

The following is an example of the YAML configuration output for Prometheus:

```yaml
targets:
  - "192.0.2.1:9090"
  - "192.0.2.2:9090"
labels:
  job: "example-service"
```

This file should be placed in your Prometheus configuration directory or included in your `prometheus.yml` under the `scrape_configs` section.

## License

`TargetGroupToPrometheus` is open-sourced under the MIT license. Feel free to clone, fork, modify, or use it in any way you see fit.
