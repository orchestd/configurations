package enviromentVariables

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"bitbucket.org/HeilaSystems/configurations/config/confgetter/utils"
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

func (e *envVariablesParamsResolver) resolveFromEnvVariables()  {
	envlist := os.Environ()
	for _, element := range envlist {
		variable := strings.Split(element, "=")
		if len(variable) == 2 {
			e.params[variable[0]] = variable[1]
		}
	}
}