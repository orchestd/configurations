package files

//import (
//	"bitbucket.org/HeilaSystems/configurations/confgetter"
//	"bitbucket.org/HeilaSystems/servicemanager/interfaces/config"
//	"encoding/json"
//	"fmt"
//	"os"
//	"strings"
//)
//
////const machineParamsFilePath = "../conf/{{HEILA_ENV}}.{{HEILA_TYPE}}.conf"
////const appParamsFilePath = "conf/app.{{HEILA_ENV}}.{{HEILA_TYPE}}.conf"
////
////type ConfFileParamsResolver struct {
////	params    config.ConfParams
////	heilaEnv  string
////	heilaType string
////}
////func NewConfParamsResolver(heilaEnv, heilaType string) config.ConfParamsResolver {
////	return &ConfFileParamsResolver{
////		heilaEnv:  heilaEnv,
////		heilaType: heilaType,
////	}
////}
////func (resolver *ConfFileParamsResolver) ResolveParams() config.ConfParams {
////	if dir, err := confgetter.GetCurrentDir(); err != nil {
////		panic(fmt.Errorf("cannot get current dir. %v", err.Error()))
////	} else {
////		if resolver.heilaEnv == "" || resolver.heilaType == "" {
////			panic(fmt.Errorf("environment variables HEILA_ENV or HEILA_TYPE are not defined"))
////		}
////		//resolve server level conf
////		serverConf := strings.Replace(machineParamsFilePath, "{{HEILA_ENV}}", resolver.heilaEnv, -1)
////		serverConf = strings.Replace(serverConf, "{{HEILA_TYPE}}", resolver.heilaType, -1)
////		serverConf = strings.ToLower(serverConf)
////
////		if loaded := resolver.resolveFromFile(dir + "/" + serverConf); loaded {
////			fmt.Println("loaded server params file")
////		}
////
////		//resolve service level conf
////		serviceConf := strings.Replace(appParamsFilePath, "{{HEILA_ENV}}", resolver.heilaEnv, -1)
////		serviceConf = strings.Replace(serviceConf, "{{HEILA_TYPE}}", resolver.heilaType, -1)
////		serviceConf = strings.ToLower(serviceConf)
////
////		if loaded := resolver.resolveFromFile(dir + "/" + serviceConf); loaded {
////			fmt.Println("loaded service params file")
////		}
////	}
////	return resolver.params
////}
////
////func (resolver *ConfFileParamsResolver) resolveFromFile(filePath string) bool {
////	fmt.Println("searching file:", filePath)
////	if _, err := os.Stat(filePath); err == nil {
////		if file, err := os.Open(filePath); err != nil {
////			panic(fmt.Errorf("while reading conf file. %v", err.Error()))
////		} else {
////			defer file.Close()
////			if err := json.NewDecoder(file).Decode(&resolver.params); err != nil {
////				panic(fmt.Errorf("cannot decode conf file to json. %v", err.Error()))
////			} else {
////				for k, v := range resolver.params {
////					if s := fmt.Sprint(v); strings.Contains(s, "{{") {
////						resolver.params[k] = resolver.params.Integrate(s)
////					}
////
////				}
////			}
////		}
////		return true
////	}
////	fmt.Println("file not found")
////	return false
////}
