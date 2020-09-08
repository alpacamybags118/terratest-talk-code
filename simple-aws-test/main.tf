terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }
}

# we are using localstack for easy local running of this example, specified by the custom endpoint.
# if you wish to deploy this to real AWS, remove the endpoints block
provider aws {
  region = var.region
}

# DynamoDB
resource aws_dynamodb_table my_dynamo_table {
  name         = var.dynamo_table_name
  billing_mode = "PROVISIONED"

  write_capacity = var.dynamo_write_capacity
  read_capacity  = var.dynamo_read_capacity

  hash_key  = var.dynamo_hash_key
  range_key = var.dynamo_range_key

  dynamic attribute {
    for_each = var.dynamo_attributes_map
    content {
      name = attribute.value["name"]
      type = attribute.value["type"]
    }
  }
}