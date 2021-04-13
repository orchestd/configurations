package enviromentVariables

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"bitbucket.org/HeilaSystems/configurations/config/confgetter/utils"
	"fmt"
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
		splittedEnv := strings.SplitN(element,"=" ,2)
		if splittedEnv[0] == "DB_HOST"{
			fmt.Print("stop")
		}
		if len(splittedEnv) == 2 {
			e.params[splittedEnv[0]] = splittedEnv[1]
		}
	}
}