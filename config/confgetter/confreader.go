package confgetter

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"bitbucket.org/HeilaSystems/configurations/config/confgetter/repos/arguments"
	"bitbucket.org/HeilaSystems/configurations/config/confgetter/repos/enviromentVariables"
	"bitbucket.org/HeilaSystems/configurations/config/confgetter/repos/files"
	"fmt"
	"reflect"
	"strings"
)

func ReadConf(conf interface{}, resolver config.ConfParamsResolver, env string) (config.Config, error) {

	argsResolver := arguments.NewArgsParamsResolver()
	argsParams := argsResolver.ResolveParams()

	envResolver := enviromentVariables.NewEnvVariablesParamsResolver()
	envParams := envResolver.ResolveParams()

	params := mergeMainAndSecondaryConfParams(argsParams, envParams)

	fileResolver := files.NewConfFileParamsResolver(env)
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
	unresolvedParams, filteredParams, _ := getAllUnresolvedParams(conf, params)
	if len(unresolvedParams) > 0 {
		panic(fmt.Errorf("found unresolved params in configuration file: " + strings.Join(unresolvedParams, ",")))
	}
	wrapper := &confgetterWrapper{instance: filteredParams}
	return wrapper, nil
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
func getAllUnresolvedParams(conf interface{}, params config.ConfParams) ([]string, ConfGetter, map[string]string) {
	val := reflect.ValueOf(conf) // could be any underlying type

	// if its a pointer, resolve its value
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	if val.Kind() != reflect.Struct {
		panic("Configuration must be a struct")
	}

	confMap := make(ConfGetter)
	typeList := make(map[string]string)
	unsolvedParams := CollectUnresolvedParams(val, params, confMap, typeList)
	return unsolvedParams, confMap, typeList
}

func CollectUnresolvedParams(val reflect.Value, params config.ConfParams, confMap ConfGetter, typeList map[string]string) []string {
	typeOfS := val.Type()
	var unsolvedParams []string
	for i := 0; i < typeOfS.NumField(); i++ {
		if val.Field(i).Kind() == reflect.Struct {
			unsolvedParams = append(unsolvedParams, CollectUnresolvedParams(val.Field(i), params, confMap, typeList)...)
			continue
		}
		var keyName string
		keyName = strings.ToLower(typeOfS.Field(i).Tag.Get("json"))
		typeName := val.Field(i).Kind().String()
		if len(keyName) == 0 {
			keyName = strings.ToLower(typeOfS.Field(i).Name)
		}
		var skip bool
		if keyNameArr := strings.Split(keyName, ","); len(keyNameArr) > 1 {
			keyName = keyNameArr[0]
			if len(keyNameArr)>= 2 && keyNameArr[1] == "omitempty"  {
				skip = true
			}

		}
		if val, ok := params[keyName]; !ok {
			if !skip{
				unsolvedParams = append(unsolvedParams, keyName)
			}
		} else {
			confMap[keyName] = val
			typeList[keyName] = typeName
		}
	}
	return unsolvedParams
}

type ConfGetter map[string]interface{}
