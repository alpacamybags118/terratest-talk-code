# AWS Global Variables

variable region {
  description = "AWS Region to create resources in"
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
  description = "Time in seconds to delay a message from being receivable after entering the queue."
  type        = number
  default     = 5
}