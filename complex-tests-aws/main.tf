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
}

module sqs {
  source = "./modules/sqs"

  region                 = var.region
  queue_name             = var.queue_name
  message_retention_time = var.message_retention_time
  delay_time             = var.delay_time
}

locals {
  env_vars = merge({
    "sqs_queue_url" : module.sqs.sqs_queue_url
  }, var.env_vars)
}

module lambda {
  source = "./modules/lambda"

  region                   = var.region
  lambda_name              = var.lambda_name
  lambda_code_path         = var.lambda_code_path
  lambda_handler           = var.lambda_handler
  lambda_runtime           = var.lambda_runtime
  lambda_memory_allocation = var.lambda_memory_allocation
  env_vars                 = local.env_vars

  sqs_queue_arn = module.sqs.sqs_queue_arn
  kms_arn       = module.sqs.kms_key_arn
}
