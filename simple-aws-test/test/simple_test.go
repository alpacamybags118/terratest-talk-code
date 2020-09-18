package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformBasicExample(t *testing.T) {

	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		VarFiles:     []string{"./example/terraform.tfvars"},
		NoColor:      true,
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)
}
