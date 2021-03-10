package credentials_test

import (
	"testing"

	"github.com/commitdev/zero/pkg/credentials"
	"github.com/stretchr/testify/assert"
)

func TestFillAWSProfileCredentials(t *testing.T) {
	mockAwsCredentialFilePath := "../../tests/test_data/aws/mock_credentials.yml"

	t.Run("fills project credentials", func(t *testing.T) {
		params := map[string]string{}
		err := credentials.FillAWSProfile(mockAwsCredentialFilePath, "default", params)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, "MOCK1_ACCESS_KEY", params["accessKeyId"])
		assert.Equal(t, "MOCK1_SECRET_ACCESS_KEY", params["secretAccessKey"])
	})

	t.Run("supports non-default profiles", func(t *testing.T) {
		params := map[string]string{}
		err := credentials.FillAWSProfile(mockAwsCredentialFilePath, "foobar", params)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, "MOCK2_ACCESS_KEY", params["accessKeyId"])
		assert.Equal(t, "MOCK2_SECRET_ACCESS_KEY", params["secretAccessKey"])
	})
}
