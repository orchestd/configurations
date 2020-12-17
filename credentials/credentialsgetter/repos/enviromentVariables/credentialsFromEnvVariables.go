package enviromentVariables

import (
	"bitbucket.org/HeilaSystems/configurations/credentials"
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type credentialsFromEnvVariables struct {
	credentials.Credentials
}

func NewCredentialsFromEnvVariables() (credentials.CredentialsGetter,error) {
	var creds credentialsFromEnvVariables
	if err := envconfig.Process("CREDS", &creds);err != nil{
		return nil, err
	}
	fmt.Println(creds)
	return &creds, nil
}

func (c *credentialsFromEnvVariables) GetCredentials() credentials.Credentials {
	return c.Credentials
}

func (c *credentialsFromEnvVariables) Implementation() interface{} {
	return c
}
