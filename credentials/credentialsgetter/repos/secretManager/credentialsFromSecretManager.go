package secretManager

import (
	"bitbucket.org/HeilaSystems/configurations/credentials"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"reflect"
	"strings"
)

type credentialsFromSecretManager struct {
	credentials.Credentials
}

func NewCredentialsFromSecretManager(projectId, version, secretName string) (credentials.CredentialsGetter, error) {
	if len(version) == 0 {
		version = "latest"
	}
	creds := credentials.Credentials{}
	//neededCreds := getSecretNamesByTag("envconfig", creds)

	c := context.Background()
	// Create the client.
	client, err := secretmanager.NewClient(c)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup client")
	}
	credsValuesMap := make(map[string]string)
	var g errgroup.Group
	g.Go(func() error {
		scrt, err := getSecretValue(projectId, secretName, version, c, client)
		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(scrt), &credsValuesMap)
		if err != nil{
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	if credsJson  , err := json.Marshal(credsValuesMap);err != nil {
		return nil, err
	}else if err := json.Unmarshal(credsJson , &creds);err != nil {
		return nil, err
	}

	return &credentialsFromSecretManager{creds}, nil
}

func (c *credentialsFromSecretManager) GetCredentials() credentials.Credentials {
	return c.Credentials
}

func (c *credentialsFromSecretManager) Implementation() interface{} {
	return c
}
func getSecretNamesByTag(tag string, s interface{}) map[string]struct{} {
	neededCreds := make(map[string]struct{})
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		panic("bad type")
	}
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get(tag), ",")[0]
		if len(v) > 0 {
			neededCreds[v] = struct{}{}
		}
	}
	return neededCreds
}

func getSecretValue(projectId, secretName string, version string, c context.Context, client *secretmanager.Client) (string, error) {

	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%v/secrets/%v/versions/%v", projectId, secretName, version),
	}
	result, err := client.AccessSecretVersion(c, accessRequest)
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %v", err)
	}
	if len(string(result.Payload.Data)) == 0 {
		return "", fmt.Errorf("secret value %s is empty", secretName)
	}
	return string(result.Payload.Data), nil
}
