# we are using localstack for easy local running of this example, specified by the custom endpoint.
# if you wish to deploy this to real AWS, remove the endpoints block
terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }
}

provider aws {
  region = var.region
  endpoints {
    sqs = "http://localhost:4566"
    kms = "http://localhost:4566"
    lambda = "http://localhost:4566"
    iam = "http://localhost:4566"
  }
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
}

module sqs {
  source = "./modules/sqs"

  region                 = var.region
  queue_name             = var.queue_name
  message_retention_time = var.message_retention_time
}

module lambda {
  source = "./module/lambda"

  region = var.region
  lambda_name = var.lambda_name
  lambda_code_path = var.lambda_code_path
  lambda_handler = var.lambda_handler
  lambda_runtime = var.lambda_runtime
  lambda_memory_allocation = var.lambda_memory_allocation
  env_vars = var.env_vars

  sqs_queue_name = module.sqs.sqs_queue_arn
}
