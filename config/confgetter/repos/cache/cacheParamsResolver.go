package cache

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"context"
)

type cacheVariablesParamsResolver struct {
	serviceName string
	env string
	params config.ConfParams
	cache config.CacheFunctions
}

func NewCacheVariablesParamsResolver(serviceName string, env string, cache config.CacheFunctions) *cacheVariablesParamsResolver {
	return &cacheVariablesParamsResolver{serviceName: serviceName, env: env, cache: cache,params: config.ConfParams{}}
}



func (e *cacheVariablesParamsResolver) ResolveParams() config.ConfParams {
	e.resolveFromCacheVariables()
	return e.params
}

func (e *cacheVariablesParamsResolver) resolveFromCacheVariables() config.ConfParams {
	if err := e.cache.GetById(context.Background(),"configurations",e.serviceName+"-"+e.env,&e.params);err != nil {
		panic(err)
	}
	return e.params
}
