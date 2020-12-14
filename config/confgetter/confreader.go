package confgetter

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"bitbucket.org/HeilaSystems/configurations/config/confgetter/repos/arguments"
	"bitbucket.org/HeilaSystems/configurations/config/confgetter/repos/enviromentVariables"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func ReadConf(conf interface{}, resolver config.ConfParamsResolver) (config.Config, error) {

	argsResolver := arguments.NewArgsParamsResolver()
	argsParams := argsResolver.ResolveParams()

	envResolver := enviromentVariables.NewEnvVariablesParamsResolver()
	envParams := envResolver.ResolveParams()

	params := mergeMainAndSecondaryConfParams(argsParams, envParams)
	if resolver == nil {
		params = mergeMainAndSecondaryConfParams(params, resolver.ResolveParams())
	}
	unresolvedParams, filteredParams := getAllUnresolvedParams(conf, params)
	if len(unresolvedParams) > 0 {
		panic(fmt.Errorf("found unresolved params in configuration file: " + strings.Join(unresolvedParams, ",")))
	}
	if marshalledParams, err := json.Marshal(filteredParams); err != nil {
		return nil, err
	} else if err := json.Unmarshal(marshalledParams, conf); err != nil {
		return nil, fmt.Errorf("cannot decode configuration map to json. %v", err.Error())
	}
	return &confgetterWrapper{instance: filteredParams}, nil
}
func mergeMainAndSecondaryConfParams(mainConfParams config.ConfParams, secondaryConfParams config.ConfParams) config.ConfParams {
	mergedConfParams := mainConfParams
	for key, val := range secondaryConfParams {
		if _, ok := mergedConfParams[key]; !ok {
			mergedConfParams[key] = val
		}
	}
	return mergedConfParams
}
func getAllUnresolvedParams(conf interface{}, params config.ConfParams) ([]string, ConfGetter) {
	confValue := reflect.ValueOf(conf)
	typeOfS := confValue.Type()
	confMap := make(ConfGetter)
	var unsolvedParams []string
	for i := 0; i < confValue.NumField(); i++ {
		var keyName string
		keyName = typeOfS.Field(i).Tag.Get("json")
		if len(keyName) == 0 {
			keyName = typeOfS.Field(i).Name
		}
		if val, ok := params[keyName]; !ok {
			unsolvedParams = append(unsolvedParams, keyName)
		} else {
			confMap[keyName] = val
		}
	}
	return unsolvedParams, confMap
}

type ConfGetter map[string]interface{}
