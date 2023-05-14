package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformS3(t *testing.T) {
	// parallel testing
	// t.Parallel()

	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	bucketID := terraform.Output(t, terraformOptions, "bucket_id")
	helloWorldMsg := terraform.Output(t, terraformOptions, "hello_world")
	//print variable to system output
	t.Log(bucketID)
	t.Log(helloWorldMsg)
	assert.Equal(t, bucketID, "my-s3-bucket")
	// output := terraform.Output(t, terraformOptions, "hello_world")
	assert.Equal(t, "hello_world", helloWorldMsg)
}