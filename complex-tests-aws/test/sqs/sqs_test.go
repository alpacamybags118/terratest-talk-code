package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestSqsModuleEnsureKMSEncryption(t *testing.T) {
	t.Parallel()

	id := random.UniqueId()
	queueName := fmt.Sprintf("sqs-queue-%s", id)

	terraformOptions := &terraform.Options{
		TerraformDir: "../../modules/sqs",
		Vars: map[string]interface{}{
			"region":     "us-east-2",
			"queue_name": queueName,
		},
		NoColor: true,
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	// Validate the queue exists and is set up for encryption
	queueURL := terraform.Output(t, terraformOptions, "sqs_queue_url")
	kmsKeyID := terraform.Output(t, terraformOptions, "kms_key_id")

	options := session.Options{
		Config: aws.Config{
			Region: aws.String("us-east-2"),
		},
	}
	session, err := session.NewSessionWithOptions(options)

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	sqsService := sqs.New(session)

	args := &sqs.GetQueueAttributesInput{
		AttributeNames: []*string{
			aws.String("KmsMasterKeyId"),
		},
		QueueUrl: aws.String(queueURL),
	}
	result, err := sqsService.GetQueueAttributes(args)

	if err != nil {
		t.Fail()
	}

	keyID := *result.Attributes["KmsMasterKeyId"]

	assert.Equal(t, kmsKeyID, keyID)

}

func TestSqsModulePutAndGetMessage(t *testing.T) {
	t.Parallel()

	id := random.UniqueId()
	queueName := fmt.Sprintf("sqs-queue-%s", id)

	terraformOptions := &terraform.Options{
		TerraformDir: "../../modules/sqs",
		Vars: map[string]interface{}{
			"region":     "us-east-2",
			"queue_name": queueName,
			"delay_time": 5,
		},
		NoColor: true,
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	// Send a message and pick it up.
	queueURL := terraform.Output(t, terraformOptions, "sqs_queue_url")

	options := session.Options{
		Config: aws.Config{
			Region: aws.String("us-east-2"),
		},
	}
	session, err := session.NewSessionWithOptions(options)

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	sqsService := sqs.New(session)

	message := &sqs.SendMessageInput{
		MessageBody: aws.String("test message"),
		QueueUrl:    aws.String(queueURL),
	}

	sqsService.SendMessage(message)

	receiveArgs := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(queueURL),
	}

	result, err := sqsService.ReceiveMessage(receiveArgs)

	if err != nil {
		t.Fail()
	}

	receivedMessage := *result.Messages[0].Body

	assert.Equal(t, "test message", receivedMessage)

}
