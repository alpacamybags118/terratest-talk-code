# we are using localstack for easy local running of this example, specified by the custom endpoint.
# if you wish to deploy this to real AWS, remove the endpoints block
provider aws {
  region = var.region
  endpoints {
    sqs = "http://localhost:4566"
  }
}

module sqs {
  source = "./modules/sqs"

  region = var.region
  queue_name = var.queue_name
  message_retention_time = var.message_retention_time
}
