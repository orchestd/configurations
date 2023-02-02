package enviromentVariables

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/orchestd/configurations/credentials"
)

type credentialsFromEnvVariables struct {
	credentials.Credentials
}

func NewCredentialsFromEnvVariables() (credentials.CredentialsGetter, error) {
	var creds credentialsFromEnvVariables
	if err := envconfig.Process("CREDS", &creds); err != nil {
		return nil, err
	}
	return &creds, nil
}

func (c *credentialsFromEnvVariables) GetCredentials() credentials.Credentials {
	return c.Credentials
}

func (c *credentialsFromEnvVariables) Implementation() interface{} {
	return c
}
