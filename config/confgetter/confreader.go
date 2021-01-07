package confgetter

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"bitbucket.org/HeilaSystems/configurations/config/confgetter/repos/arguments"
	"bitbucket.org/HeilaSystems/configurations/config/confgetter/repos/enviromentVariables"
	"bitbucket.org/HeilaSystems/configurations/config/confgetter/repos/files"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func ReadConf(conf interface{}, resolver config.ConfParamsResolver,env string) (config.Config, error) {

	argsResolver := arguments.NewArgsParamsResolver()
	argsParams := argsResolver.ResolveParams()

	envResolver := enviromentVariables.NewEnvVariablesParamsResolver()
	envParams := envResolver.ResolveParams()

	params := mergeMainAndSecondaryConfParams(argsParams, envParams)

	fileResolver  := files.NewConfFileParamsResolver(env)
	filesParams := fileResolver.ResolveParams()
	if filesParams != nil {
		params = mergeMainAndSecondaryConfParams(params, filesParams)
	}

	if resolver == nil {
		resolverParams := resolver.ResolveParams()
		if resolverParams != nil {
			params = mergeMainAndSecondaryConfParams(params, resolverParams)
		}
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
	val := reflect.ValueOf(conf) // could be any underlying type

	// if its a pointer, resolve its value
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	if val.Kind() != reflect.Struct {
		panic("Configuration must be a struct")
	}


	confMap := make(ConfGetter)
	var unsolvedParams []string
	CollectUnresolvedParams(val,params,confMap,unsolvedParams)
	return unsolvedParams, confMap
}

func CollectUnresolvedParams(val reflect.Value , params config.ConfParams ,confMap ConfGetter , unsolvedParams []string) {
	typeOfS := val.Type()
	for i := 0; i < typeOfS.NumField(); i++ {
		if val.Field(i).Kind() == reflect.Struct {
			CollectUnresolvedParams(val.Field(i) ,params,confMap,unsolvedParams )
			continue;
		}
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
}

type ConfGetter map[string]interface{}
