package confgetter

import (
	"container/list"
	"fmt"
	"github.com/orchestd/configurations/config"
)

type confReaderConfig struct {
	Env         *string
	ServiceName *string
	Resolver    config.ConfParamsResolver
	Conf        interface{}
}

type defaultConfReaderBuilder struct {
	ll *list.List
}

func Builder() config.Builder {
	return &defaultConfReaderBuilder{ll: list.New()}
}

func (cr *defaultConfReaderBuilder) SetEnv(env string) config.Builder {
	cr.ll.PushBack(func(cfg *confReaderConfig) {
		cfg.Env = &env
	})
	return cr
}
func (cr *defaultConfReaderBuilder) SetServiceName(name string) config.Builder {
	cr.ll.PushBack(func(cfg *confReaderConfig) {
		cfg.ServiceName = &name
	})
	return cr
}

func (cr *defaultConfReaderBuilder) SetRepo(resolver config.ConfParamsResolver) config.Builder {
	cr.ll.PushBack(func(cfg *confReaderConfig) {
		cfg.Resolver = resolver
	})
	return cr
}
func (cr *defaultConfReaderBuilder) SetConfStruct(conf interface{}) config.Builder {
	cr.ll.PushBack(func(cfg *confReaderConfig) {
		cfg.Conf = conf
	})
	return cr
}
func (cr *defaultConfReaderBuilder) Build() (config.Config, error) {
	confreaderCfg := &confReaderConfig{}
	for e := cr.ll.Front(); e != nil; e = e.Next() {
		f := e.Value.(func(cfg *confReaderConfig))
		f(confreaderCfg)
	}
	if confreaderCfg.Env == nil {
		return nil, fmt.Errorf("cannot initalize configurations without env settings")
	} else if confreaderCfg.ServiceName == nil {
		return nil, fmt.Errorf("cannot initalize configurations without service name")
	} /*else if confreaderCfg.Resolver == nil {
		return nil, fmt.Errorf("cannot initalize configurations without repo")
	}*/
	return ReadConf(confreaderCfg.Conf, confreaderCfg.Resolver, *confreaderCfg.Env)
}

var _ config.Builder = (*defaultConfReaderBuilder)(nil)
