package credentialsgetter

import (
	"bitbucket.org/HeilaSystems/configurations/credentials/credentialsgetter/repos/enviromentVariables"
	"bitbucket.org/HeilaSystems/configurations/credentials"
	"container/list"
)

type credentialsConfig struct {
	devMode bool
}

type defaultCredentialsConfigBuilder struct {
	ll *list.List
}

func (cr *defaultCredentialsConfigBuilder) DevMode() credentials.Builder {
	cr.ll.PushBack(func(cfg *credentialsConfig) {
		cfg.devMode = true
	})
	return cr
}

func (cr *defaultCredentialsConfigBuilder) Build() (credentials.CredentialsGetter, error) {
	credsCnf := &credentialsConfig{}
	for e := cr.ll.Front(); e != nil; e = e.Next() {
		f := e.Value.(func(cfg *credentialsConfig))
		f(credsCnf)
	}

	if credsCnf.devMode {
		return enviromentVariables.NewCredentialsFromEnvVariables()
	} else {
		return enviromentVariables.NewCredentialsFromEnvVariables()
	}
}

func Builder() credentials.Builder {
	return &defaultCredentialsConfigBuilder{ll: list.New()}
}
var _ credentials.Builder = (*defaultCredentialsConfigBuilder)(nil)
