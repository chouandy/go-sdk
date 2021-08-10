package dotenv

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var secretsPassword string

var (
	secretsPasswordPsName   string
	secretsPasswordPsRegion string
)

// GetSecretsPassword get secrects password
func GetSecretsPassword() string {
	// Password from env
	secretsPassword = os.Getenv("SECRETS_PASSWORD")
	if len(secretsPassword) > 0 {
		return secretsPassword
	}

	// Get parameter store envs
	secretsPasswordPsName = os.Getenv("SECRETS_PASSWORD_PS_NAME")
	secretsPasswordPsRegion = os.Getenv("SECRETS_PASSWORD_PS_REGION")
	// Check parameter store envs
	if len(secretsPasswordPsName) > 0 && len(secretsPasswordPsRegion) > 0 {
		// Get password from aws parameter store
		password, err := GetSecretsPasswordFromParameterStore()
		if err != nil {
			fmt.Printf("[GetSecretsPassword] [GetSecretsPasswordFromParameterStore] %s\n", err)
			return ""
		}

		return password
	}

	return ""
}

// GetSecretsPasswordFromParameterStore get secrets password fromparameterstore
func GetSecretsPasswordFromParameterStore() (string, error) {
	// new aws ssm client
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}
	cfg.Region = secretsPasswordPsRegion
	client := ssm.NewFromConfig(cfg)

	// New input
	input := &ssm.GetParameterInput{
		Name:           aws.String(secretsPasswordPsName),
		WithDecryption: true,
	}

	// New request
	resp, err := client.GetParameter(context.TODO(), input)
	if err != nil {
		return "", err
	}

	return aws.ToString(resp.Parameter.Value), nil
}
