# TargetGroupToPrometheus

`TargetGroupToPrometheus` is a CLI tool developed in Go, designed to generate Prometheus target configurations from AWS Elastic Load Balancer (ELB) target groups. 

It retrieves public IP addresses of instances registered with a specified target group and generates a YAML configuration file for Prometheus, facilitating dynamic service discovery.

## Pre-requisites

- Go 1.20 or newer.
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

### AWS Credentials Configuration

It is recommended to set up your AWS credentials using the file-based configuration method. This involves creating or updating the AWS credentials file typically located at `~/.aws/credentials` on Linux and macOS, or `%USERPROFILE%\.aws\credentials` on Windows. This file should contain your access key ID and secret access key under a profile name.

By default, the tool uses the AWS credentials stored under the `[default]` profile in this file. If you wish to use a different profile, you can specify this using the `-aws-profile` flag when running the tool. This allows you to easily switch between different AWS accounts or configurations.

Example of a credentials file:

```
[default]
aws_access_key_id = YOUR_DEFAULT_ACCESS_KEY_ID
aws_secret_access_key = YOUR_DEFAULT_SECRET_ACCESS_KEY

[another-profile]
aws_access_key_id = YOUR_OTHER_ACCESS_KEY_ID
aws_secret_access_key = YOUR_OTHER_SECRET_ACCESS_KEY
```

To use a profile other than the default, run the tool with the `-aws-profile` flag, like so:

```sh
./tg2prom.exe -aws-profile another-profile -target-group-arn "your-target-group-arn" -listening-port 8080 -job-name "your-job-name" -output "your-output-file.yaml"
```

If you prefer not to use the file-based method or are running the tool in an environment where setting up a credentials file is not feasible, you can directly provide AWS credentials via the `-aws-access-key-id` and `-aws-secret-access-key` flags. However, please be aware of the security implications of passing sensitive information through command-line arguments.

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

This file should be placed in your Prometheus configuration directory or included in your `prometheus.yml` under the `scrape_configs` section, in example like that:

```sh
scrape_configs:
 - job_name: 'your-job-name'
   file_sd_configs:
    - files:
       - 'your-output-file.yaml'
```

## License

`TargetGroupToPrometheus` is open-sourced under the MIT license. Feel free to clone, fork, modify, or use it in any way you see fit.
