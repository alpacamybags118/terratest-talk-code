{
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": [
          "sqs:"
        ],
        "Effect": "Allow",
        "Resource": ${sqs_queue_arn}
      }
    ]
  }