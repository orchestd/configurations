package enviromentVariables

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/config"
	"os"
	"strings"
)

type envVariablesParamsResolver struct {
	params config.ConfParams
}

func NewEnvVariablesParamsResolver() config.ConfParamsResolver {
	return &envVariablesParamsResolver{params: config.ConfParams{}}
}

func (e *envVariablesParamsResolver) ResolveParams() config.ConfParams {
	e.resolveFromEnvVariables()
	return e.params
}

func (e *envVariablesParamsResolver) resolveFromEnvVariables()  {
	envlist := os.Environ()
	for _, element := range envlist {
		variable := strings.Split(element, "=")
		if len(variable) == 2 {
			e.params[variable[0]] = variable[1]
		}
	}
}