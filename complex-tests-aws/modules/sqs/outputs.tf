output "sqs_queue_url" {
  value = aws_sqs_queue.app_queue.id
}

output "kms_key_id" {
  value = aws_kms_key.sqs_kms.key_id
}

output "sqs_queue_arn" {
  value = aws_sqs_queue.app_queue.arn
}

output "sqs_queue_name" {
  value = aws_sqs_queue.app_queue.name
}

output "kms_key_arn" {
  value = aws_kms_key.sqs_kms.arn
}