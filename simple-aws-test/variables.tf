# AWS Global Variables

variable region {
  description = "AWS Region to create resources in"
  type        = string
}


# DynamoDB Variables

variable dynamo_table_name {
  description = "Name to give dynamo table"
}

variable dynamo_write_capacity {
  description = "Write capacity units for the table. Defaults to 10"
  type        = number
  default     = 10
}

variable dynamo_read_capacity {
  description = "Read capacity units for the table. Defaults to 10"
  type        = number
  default     = 10
}

variable dynamo_hash_key {
  description = "The range key for your dynamo table. This will need to be provided in the attributes variable map."
  type        = string
}

variable dynamo_range_key {
  description = "The optional range (sort) for your dynamo table. If specififed, this will need to be provided in the attributes variable map."
  type        = string
  default     = null
}

variable dynamo_attributes_map {
  description = <<EOF
    "The map of all attributes of your table. Each entry contains the name and data type. Must provide attribute for hash key and range key (if provided).
    Example: [
      {
        name: "name"
        type: "S"
      },
      {
        name: "type",
        type: "S"
      }
    ]
    "
    EOF

  type = list(map(string))
  validation {
    condition     = length(var.dynamo_attributes_map) >= 1
    error_message = "You must have at least one element in your attributes map."
  }
}