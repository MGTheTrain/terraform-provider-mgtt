package mgtt

import (
	"os"
	"testing"

	"github.com/MGTheTrain/terraform-provider-mgtt/mgtt"
	"github.com/stretchr/testify/assert"
)

func TestAwsS3BucketHandler(t *testing.T) {
	// Read parameters from environment variables
	accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	bucketName := "test-bucket-12345"
	region := "us-west-2"

	if accessKeyId == "" || secretAccessKey == "" {
		t.Fatal("Missing required environment variables")
	}

	awsS3BucketHandler, err := mgtt.NewAwsS3BucketHandler(accessKeyId, secretAccessKey, bucketName, region)
	assert.NoError(t, err, "NewAwsS3BucketHandler should not return an error")

	// [C]reate
	err = awsS3BucketHandler.CreateAwsS3Bucket()
	assert.NoError(t, err, "CreateAwsS3Bucket should not return an error")

	// [D]elete
	awsS3BucketHandler.DeleteAwsS3Bucket()
	assert.NoError(t, err, "DeleteAwsS3Bucket should not return an error")

}
