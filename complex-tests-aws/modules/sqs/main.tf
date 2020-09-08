provider aws {
  region = var.region
}

resource "aws_kms_key" "sqs_kms" {
  description             = "KMS key for SQS"
  deletion_window_in_days = 10
}

resource "aws_sqs_queue" "app_queue" {
  name                      = var.queue_name
  delay_seconds             = var.delay_time
  max_message_size          = 262144
  message_retention_seconds = var.message_retention_time
  receive_wait_time_seconds = 10
  kms_master_key_id         = aws_kms_key.sqs_kms.id

  tags = {
    Environment = "production"
  }
}
