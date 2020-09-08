# complex-tests-aws

A complex example consisting of multiple terraform modules, and tests for each. This module create the following:

1. SQS Queue
2. Lambda function
3. IAM role and permissions for lambda function to write to queue

You will need terraform 0.13 to run this example.

## Running Terraform

You can run `terraform plan --var-file=example/terraform.tfvars` to validate what will be created and `terraform apply --var-file=example/terraform.tfvars` using the example data.

## Running Tests

You will need to have `go` installed in order to run the test.
To run tests, follow these steps:

1. Run `cd test`
2. Run `go mod init complex-tests-aws`
3. Run `go test -v -timeout 30m`