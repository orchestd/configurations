 package arguments

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"os"
	"strings"
)

type argsParamsResolver struct {
	params config.ConfParams
}

func NewArgsParamsResolver() config.ConfParamsResolver {
	return &argsParamsResolver{params: config.ConfParams{}}
}

func (e *argsParamsResolver) ResolveParams() config.ConfParams {
	e.resolveFromOsArgs()
	return e.params
}

func (e *argsParamsResolver) resolveFromOsArgs()  {
	envlist := os.Args
	for _, element := range envlist {
		variable := strings.Split(element, "=")
		if len(variable) == 2 {
			e.params[variable[0]] = variable[1]
		}
	}
}
