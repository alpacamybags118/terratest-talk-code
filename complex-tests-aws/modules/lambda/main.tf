provider aws {
  region = var.region
  endpoints {
    lambda = "http://localhost:4566"
    iam = "http://localhost:4566"
  }
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
}


resource "aws_iam_role" "lambda_iam_role" {
  name = "${var.lambda_name}-iam-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_function" "lambda" {
  filename      = var.lambda_code_path
  function_name = var.lambda_name
  role          = aws_iam_role.lambda_iam_role.arn
  handler       = var.lambda_handler

  source_code_hash = filebase64sha256(var.lambda_code_path)

  runtime = var.lambda_runtime

  dynamic environment {
    for_each = var.env_vars != null ? [1] : []
    content {
      variables = var.env_vars
    }
  }
}

# This is commented out on purpose to show a failing test due to missing permission. To "fix" the test issue, uncomment this code.
/*
resource "aws_iam_policy" "sqs_policy" {
  count = var.sqs_queue_name ? 1 : 0

  name        = "${var.lambda_name}-sqs-access-policy"
  description = "Policy to give lambda access to SQS queue"

  policy = templatefile("${path.module}/files/policy.tpl", {sqs_queue_arn = var.sqs_queue_arn})

resource "aws_iam_role_policy_attachment" "sqs_policy_attach" {
  var.sqs_queue_name ? 1 : 0

  role       = aws_iam_role.lambda_iam_role.name
  policy_arn = aws_iam_policy.sqs_policy[0].arn
}
*/