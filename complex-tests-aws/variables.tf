# AWS Global Variables

variable region {
  description = "AWS Regions to create resources in"
  type        = string
}


# SQS
variable queue_name {
  description = "Name to give the SQS queue"
  type        = string

  validation {
    condition     = element(concat(regexall("[[:alnum:]_-]{1,80}", var.queue_name), [""]), 0) == var.queue_name
    error_message = "Your queue name must be 80 character or less and can contain alphanumeric characters, hyphens (- ), and underscores (_ )."
  }
}

variable message_retention_time {
  description = "Time, in seconds, to retain a message in the queue. Default is 4 days."
  type        = number
  default     = 345600

  validation {
    condition     = var.message_retention_time >= 60 && var.message_retention_time <= 1209600
    error_message = "Message retention time must be between 60 and 1209600 seconds."
  }
}


variable delay_time {
  description = ""
  type        = number
  default     = 5
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