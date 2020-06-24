package credentials_test

import (
	"testing"

	"github.com/commitdev/zero/internal/config/globalconfig"
	"github.com/commitdev/zero/pkg/credentials"
	"github.com/stretchr/testify/assert"
)

func TestFillAWSProfileCredentials(t *testing.T) {
	mockAwsCredentialFilePath := "../../tests/test_data/aws/mock_credentials.yml"
	t.Run("fills project credentials", func(t *testing.T) {
		projectCreds := globalconfig.ProjectCredential{}
		projectCreds = credentials.GetAWSProfileCredentials(mockAwsCredentialFilePath, "default", projectCreds)
		assert.Equal(t, "MOCK1_ACCESS_KEY", projectCreds.AWSResourceConfig.AccessKeyID)
		assert.Equal(t, "MOCK1_SECRET_ACCESS_KEY", projectCreds.AWSResourceConfig.SecretAccessKey)
	})

	t.Run("supports non-default profiles", func(t *testing.T) {
		projectCreds := globalconfig.ProjectCredential{}
		projectCreds = credentials.GetAWSProfileCredentials(mockAwsCredentialFilePath, "foobar", projectCreds)
		assert.Equal(t, "MOCK2_ACCESS_KEY", projectCreds.AWSResourceConfig.AccessKeyID)
		assert.Equal(t, "MOCK2_SECRET_ACCESS_KEY", projectCreds.AWSResourceConfig.SecretAccessKey)
	})
}
