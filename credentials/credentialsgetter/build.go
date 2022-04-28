package credentialsgetter

import (
	"bitbucket.org/HeilaSystems/configurations/credentials"
	"bitbucket.org/HeilaSystems/configurations/credentials/credentialsgetter/repos/enviromentVariables"
	"bitbucket.org/HeilaSystems/configurations/credentials/credentialsgetter/repos/secretManager"
	"container/list"
	"fmt"
)

type credentialsConfig struct {
	useSecretManager bool
	projectId        string
	overrideVersion  string
	secretName       string
}

type defaultCredentialsConfigBuilder struct {
	ll *list.List
}

func (cr *defaultCredentialsConfigBuilder) UseGcpSecretManager(projectId string) credentials.Builder {
	cr.ll.PushBack(func(cfg *credentialsConfig) {
		cfg.useSecretManager = true
		cfg.projectId = projectId
	})
	return cr
}

func (cr *defaultCredentialsConfigBuilder) SetSecretName(name string) credentials.Builder {
	cr.ll.PushBack(func(cfg *credentialsConfig) {
		cfg.secretName = name
	})
	return cr
}

func (cr *defaultCredentialsConfigBuilder) SetSecretManagerVersion(version string) credentials.Builder {
	cr.ll.PushBack(func(cfg *credentialsConfig) {
		cfg.overrideVersion = version
	})
	return cr
}

func (cr *defaultCredentialsConfigBuilder) Build() (credentials.CredentialsGetter, error) {
	credsCnf := &credentialsConfig{}
	for e := cr.ll.Front(); e != nil; e = e.Next() {
		f := e.Value.(func(cfg *credentialsConfig))
		f(credsCnf)
	}

	if !credsCnf.useSecretManager {
		return enviromentVariables.NewCredentialsFromEnvVariables()
	} else if len(credsCnf.projectId) == 0 {
		return nil, fmt.Errorf("Cannot initialize secret manager as credentials repo without PROJECT_ID ")
	} else {
		return secretManager.NewCredentialsFromSecretManager(credsCnf.projectId, credsCnf.overrideVersion, credsCnf.secretName)
	}
}

func Builder() credentials.Builder {
	return &defaultCredentialsConfigBuilder{ll: list.New()}
}

var _ credentials.Builder = (*defaultCredentialsConfigBuilder)(nil)
