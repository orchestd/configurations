package enviromentVariables

import (
	"github.com/orchestd/configurations/config"
	"github.com/orchestd/configurations/config/confgetter/utils"
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
	return utils.MapToLowercaseMapToLowercase(e.params)
}

func (e *envVariablesParamsResolver) resolveFromEnvVariables() {
	envlist := os.Environ()
	for _, element := range envlist {
		if splittedEnv := strings.SplitN(element, "=", 2); len(splittedEnv) == 2 {
			e.params[splittedEnv[0]] = splittedEnv[1]
		}
	}
}
