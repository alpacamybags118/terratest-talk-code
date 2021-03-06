package test

import (
	"fmt"
	"testing"

	"github.com/TwinProduction/go-color"

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
		Payload:      []byte(fmt.Sprintf("{\"sqs_queue_url\": \"%s\"}", queueURL)),
	}
	result, err := lambdaService.Invoke(args)

	if err != nil {
		fmt.Println(color.Ize(color.Red, err.Error()))
		t.FailNow()
	}

	var functionError string
	if result.FunctionError == nil {
		functionError = ""
	} else {
		functionError = *result.FunctionError
	}

	payload := string(result.Payload)

	if len(functionError) > 0 {
		fmt.Println(color.Ize(color.Red, "Lambda execution failed with following error:"))
		fmt.Println(color.Ize(color.Red, payload))
		t.FailNow()
	}

	fmt.Println(color.Ize(color.Yellow, "Lambda executed successfully, picking up messsage from queue."))

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

	receivedMessage := *sqsResult.Messages[0].Body

	fmt.Println(color.Ize(color.Yellow, "Sent message: Hello, world"))
	fmt.Println(color.Ize(color.Yellow, fmt.Sprintf("Received message: %s", receivedMessage)))

	assert.Equal(t, "Hello, world", receivedMessage)

}
