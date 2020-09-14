# complex-tests-aws

A complex example consisting of multiple terraform modules, and tests for each. This module create the following:

1. SQS Queue
2. Lambda function
3. IAM role and permissions for lambda function to write to queue

You will need terraform 0.13 to run this example.

## Running Terraform

You can run `terraform plan --var-file=example/terraform.tfvars` to validate what will be created and `terraform apply --var-file=example/terraform.tfvars` using the example data.

## Simulate Test Failure

This example has a means to easily create a failure scenerio for the module test. Comment out the following code in the `lambda` module `main.tf`

```
resource "aws_iam_policy" "sqs_policy" {
  name        = "${var.lambda_name}-sqs-access-policy"
  description = "Policy to give lambda access to SQS queue"

  policy = templatefile("${path.module}/files/policy.tpl", { sqs_queue_arn = var.sqs_queue_arn, kms_arn = var.kms_arn })
}

resource "aws_iam_role_policy_attachment" "sqs_policy_attach" {
  role       = aws_iam_role.lambda_iam_role.name
  policy_arn = aws_iam_policy.sqs_policy.arn
}

```

## Running Tests

You will need to have `go` installed in order to run the test.
To run tests, follow these steps:

1. Run `cd test`
2. Run `go mod init complex-tests-aws`
3. Run `go test -v -timeout 30m`