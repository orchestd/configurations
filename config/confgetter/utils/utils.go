package utils

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"strings"
)

func MapToLowercaseMapToLowercase(params config.ConfParams) config.ConfParams {
	lowerMap := make(config.ConfParams)
	for s, v := range params {
		lowerMap[strings.ToLower(s)] = v
	}
	return lowerMap
}
