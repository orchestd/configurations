package files

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"bitbucket.org/HeilaSystems/configurations/config/confgetter/utils"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const appParamsFilePath = "conf/{{HEILA_ENV}}/app.conf"

type ConfFileParamsResolver struct {
	params config.ConfParams

	env string
}

func NewConfFileParamsResolver(env string) config.ConfParamsResolver {
	return &ConfFileParamsResolver{
		env: env,
	}
}
func (resolver *ConfFileParamsResolver) ResolveParams() config.ConfParams {
	if err, dir := GetExecutableDir(); err != nil {
		panic(fmt.Errorf("cannot get current dir. %v", err.Error()))
	} else {
		if resolver.env == "" {
			panic(fmt.Errorf("environment variables HEILA_ENV is not defined"))
		}
		//resolve service level conf
		serviceConf := strings.Replace(appParamsFilePath, "{{HEILA_ENV}}", resolver.env, -1)
		serviceConf = strings.ToLower(serviceConf)

		_ = resolver.resolveFromFile(dir + "/" + serviceConf)

	}
	return utils.MapToLowercaseMapToLowercase(resolver.params)
}

func (resolver *ConfFileParamsResolver) resolveFromFile(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		if file, err := os.Open(filePath); err != nil {
			panic(fmt.Errorf("while reading conf file. %v", err.Error()))
		} else {
			defer file.Close()
			if err := json.NewDecoder(file).Decode(&resolver.params); err != nil {
				panic(fmt.Errorf("cannot decode conf file to json. %v", err.Error()))
			}
		}
		return true
	}
	fmt.Printf("conf file %s Not Found \n",filePath)
	return false
}
func GetExecutableDir() (error, string) {
	ex, err := os.Executable()
	if err != nil {
		return err, ""
	}
	exPath := filepath.Dir(ex)
	return nil, exPath
}
