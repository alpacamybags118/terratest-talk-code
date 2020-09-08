output "queue_url" {
  value = module.sqs.sqs_queue_url
}

output "lambda_arn" {
  value = module.lambda.lambda_arn
}