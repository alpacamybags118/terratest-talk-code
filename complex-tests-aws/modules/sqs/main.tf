# SQS
resource "aws_sqs_queue" "app_queue" {
  name                      = var.queue_name
  delay_seconds             = 90
  max_message_size          = 2048
  message_retention_seconds = var.message_retention_time
  receive_wait_time_seconds = 10

  tags = {
    Environment = "production"
  }
}
