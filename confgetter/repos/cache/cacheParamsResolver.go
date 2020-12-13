package cache

import (
	"bitbucket.org/HeilaSystems/dependencybundler/interfaces/config"
)

type cacheVariablesParamsResolver struct {
	params config.ConfParams
}

func NewCacheVariablesParamsResolver() *cacheVariablesParamsResolver {
	return &cacheVariablesParamsResolver{params: config.ConfParams{}}
}

func (e *cacheVariablesParamsResolver) ResolveParams() config.ConfParams {
	e.resolveFromCacheVariables()
	return e.params
}

func (e *cacheVariablesParamsResolver) resolveFromCacheVariables() config.ConfParams {
	return e.params
}
