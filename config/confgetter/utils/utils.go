package utils

import (
	"github.com/orchestd/configurations/config"
	"strings"
)

func MapToLowercaseMapToLowercase(params config.ConfParams) config.ConfParams {
	lowerMap := make(config.ConfParams)
	for s, v := range params {
		lowerMap[strings.ToLower(s)] = v
	}
	return lowerMap
}
