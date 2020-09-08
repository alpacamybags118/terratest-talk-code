package test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestModuleE2E(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../../",
		VarFiles:     []string{"./example/terraform.tfvars"},
		NoColor:      true,
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	// step 1 invoke the lambda to put a messsage in the queue
	lambdaARN := terraform.Output(t, terraformOptions, "lambda_arn")
	queueURL := terraform.Output(t, terraformOptions, "queue_url")

	options := session.Options{
		Config: aws.Config{
			Region: aws.String("us-east-2"),
		},
	}
	session, err := session.NewSessionWithOptions(options)

	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	lambdaService := lambda.New(session)

	args := &lambda.InvokeInput{
		FunctionName: aws.String(lambdaARN),
	}
	result, err := lambdaService.Invoke(args)

	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	status := *result.StatusCode

	assert.Equal(t, int64(200), status)

	// step 2 pick up the message from the queue
	sqsService := sqs.New(session)

	receiveArgs := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(queueURL),
	}

	sqsResult, err := sqsService.ReceiveMessage(receiveArgs)

	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(sqsResult)
	receivedMessage := *sqsResult.Messages[0].Body

	assert.Equal(t, "Hello, world", receivedMessage)

}
