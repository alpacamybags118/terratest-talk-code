# AWS Global Variables

variable region {
  description = "AWS Region to create resources in"
  type        = string
}


# Lambda
variable lambda_name {
  description = "Name to give lambda function"
  type        = string
}

variable lambda_code_path {
  description = "Path to find lambda code. e.g.: code.zip"
  type        = string
}

variable lambda_handler {
  description = "Name of handler in code that is the entrypoint for the lambda"
  type        = string
}

variable lambda_runtime {
  description = "Runtime of lambda. See here for a list of runtimes: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html"
  type        = string
}

variable lambda_memory_allocation {
  description = "Memory to allocate to lambda. This proportionally impacts the CPU allocation to the lambda as well"
  type        = number
  default     = 128
}

variable env_vars {
  description = <<EOL
    Environment variables to provide to lambda
    e.g: {
        foo = "bar"
    }
    EOL
  type        = map(string)
  default     = null
}

# SQS
variable sqs_queue_arn {
  description = "ARN of SQS queue to grant lambda permissions to"
  type        = string
  default     = null
}

# KMS
variable kms_arn {
  description = "ARN of KMS key used to encrypt messages in SQS"
  type        = string
  default     = null
}