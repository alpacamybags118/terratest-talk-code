const AWS = require('aws-sdk');
const sqs = new AWS.SQS({region: 'us-east-2'});

exports.handler =  async function(event, context) {
    const queueUrl = event.sqs_queue_url;

    const sendMessageArgs = {
        MessageBody: "Hello, world",
        QueueUrl: queueUrl
    }
    const response = await sqs.sendMessage(sendMessageArgs).promise()

    return response;
  }