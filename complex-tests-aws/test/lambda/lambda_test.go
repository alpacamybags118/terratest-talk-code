package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestLambdaInvoke(t *testing.T) {
	t.Parallel()

	id := random.UniqueId()
	lambdaName := fmt.Sprintf("lambda-%s", id)

	terraformOptions := &terraform.Options{
		TerraformDir: "../../modules/lambda",
		Vars: map[string]interface{}{
			"region":           "us-east-2",
			"lambda_name":      lambdaName,
			"lambda_code_path": "../../test/lambda/lambda_test.zip",
			"lambda_handler":   "index.handler",
			"lambda_runtime":   "nodejs12.x",
		},
		NoColor: true,
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	// Invoke the lambda and ensure an expected response
	lambdaARN := terraform.Output(t, terraformOptions, "lambda_arn")

	options := session.Options{
		Config: aws.Config{
			Endpoint: aws.String("http://localhost:4566"),
			Region:   aws.String("us-east-2"),
		},
	}
	session, err := session.NewSessionWithOptions(options)

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	lambdaService := lambda.New(session)

	args := &lambda.InvokeInput{
		FunctionName: aws.String(lambdaARN),
	}
	result, err := lambdaService.Invoke(args)

	if err != nil {
		t.Fail()
	}

	status := *result.StatusCode
	response := strings.TrimSuffix(string(result.Payload), "\n")

	assert.Equal(t, int64(200), status)
	assert.Equal(t, `"Hello, world!"`, response)
}
