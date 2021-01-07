package cache

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/cache"
	"context"
)

type cacheVariablesParamsResolver struct {
	serviceName string
	env string
	version string
	params config.ConfParams
	cache cache.CacheStorageGetter
}

func NewCacheVariablesParamsResolver(serviceName string, env string, version string,cache cache.CacheStorageGetter) *cacheVariablesParamsResolver {
	return &cacheVariablesParamsResolver{serviceName: serviceName, env: env,version: version,cache: cache,params: config.ConfParams{}}
}



func (e *cacheVariablesParamsResolver) ResolveParams() config.ConfParams {
	e.resolveFromCacheVariables()
	return e.params
}

func (e *cacheVariablesParamsResolver) resolveFromCacheVariables() config.ConfParams {
	if err := e.cache.GetById(context.Background(),"configurations",e.version,e.serviceName+"-"+e.env,&e.params);err != nil {
		panic(err)
	}
	return e.params
}
