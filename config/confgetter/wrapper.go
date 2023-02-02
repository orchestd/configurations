package confgetter

import (
	"encoding/json"
	"fmt"
	"github.com/orchestd/configurations/config"
	"github.com/orchestd/sharedlib/consts"
	"github.com/spf13/cast"
	"strings"
	"time"
)

func NewConfgetterWrapper(instance map[string]interface{}) confgetterWrapper {
	return confgetterWrapper{
		instance: instance,
	}
}

type confgetterWrapper struct {
	instance ConfGetter
}

func (v *confgetterWrapper) Get(key string) config.Value {
	return &valueWrapper{
		key:               strings.ToLower(key),
		confgetterWrapper: v,
	}
}

func (v *confgetterWrapper) GetServiceName() (string, error) {
	return v.Get(consts.ServiceNameEnv).String()
}

func (v *confgetterWrapper) Implementation() interface{} {
	return v.instance
}

func (v *confgetterWrapper) ShouldRunDebug(action string, debug bool) (bool, error) {
	// TODO: add action map to follow/allow debug actions(=flows)
	if !debug {
		return false, nil
	} else if env, err := v.Get(consts.HeilaEnv).String(); err != nil {
		return false, err
	} else {
		return env != consts.EnvProd, nil
	}
}

type valueWrapper struct {
	key string
	*confgetterWrapper
}

func (v *valueWrapper) IsSet() bool {
	_, ok := v.instance[v.key]
	return ok
}

func (v *valueWrapper) Raw() (interface{}, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return nil, KeyNotFoundError(v.key)
	}
	return val, nil
}

func (v *valueWrapper) Bool() (bool, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return false, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToBoolE(val); err != nil {
		return false, CastingError(err, v.key, val, "bool")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) Int() (int, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToIntE(val); err != nil {
		return 0, CastingError(err, v.key, val, "int")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) Int32() (int32, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToInt32E(val); err != nil {
		return 0, CastingError(err, v.key, val, "int32")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) Int64() (int64, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToInt64E(val); err != nil {
		return 0, CastingError(err, v.key, val, "int64")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) Uint() (uint, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToUintE(val); err != nil {
		return 0, CastingError(err, v.key, val, "uint")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) Uint32() (uint32, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToUint32E(val); err != nil {
		return 0, CastingError(err, v.key, val, "uint32")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) Uint64() (uint64, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToUint64E(val); err != nil {
		return 0, CastingError(err, v.key, val, "uint64")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) Float64() (float64, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToFloat64E(val); err != nil {
		return 0, CastingError(err, v.key, val, "float64")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) Time() (time.Time, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return time.Time{}, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToTimeE(val); err != nil {
		return time.Time{}, CastingError(err, v.key, val, "time.Time")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) Duration() (time.Duration, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToDurationE(val); err != nil {
		return 0, CastingError(err, v.key, val, "time.Duration")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) String() (string, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return "", KeyNotFoundError(v.key)
	}
	if res, err := cast.ToStringE(val); err != nil {
		return "", CastingError(err, v.key, val, "string")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) IntSlice() ([]int, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return nil, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToIntSliceE(val); err != nil {
		return nil, CastingError(err, v.key, val, "[]int")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) StringSlice() ([]string, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return nil, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToStringSliceE(val); err != nil {
		return nil, CastingError(err, v.key, val, "[]string")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) StringMap() (map[string]interface{}, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return nil, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToStringMapE(val); err != nil {
		return nil, CastingError(err, v.key, val, "map[string]interface{}")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) StringMapString() (map[string]string, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return nil, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToStringMapStringE(val); err != nil {
		return nil, CastingError(err, v.key, val, "map[string]string")
	} else {
		return res, nil
	}
}

func (v *valueWrapper) StringMapStringSlice() (map[string][]string, error) {
	val, ok := v.instance[v.key]
	if !ok {
		return nil, KeyNotFoundError(v.key)
	}
	if res, err := cast.ToStringMapStringSliceE(val); err != nil {
		return nil, CastingError(err, v.key, val, "map[string][]string")
	} else {
		return res, nil
	}
}

// Unmarshal uses default decoder options, if you need some special behavior than it's best to get cfg.Implementation() and use it from there
func (v *valueWrapper) Unmarshal(result interface{}) error {
	val, ok := v.instance[v.key]
	if !ok {
		return KeyNotFoundError(v.key)
	}
	if byteArray, err := json.Marshal(val); err != nil {
		return err
	} else {
		return json.Unmarshal(byteArray, &result)
	}
}

func CastingError(err error, key string, rawVal interface{}, confType string) error {
	return fmt.Errorf(`Cannot parse %s with value "%v" into %s. %v`, key, rawVal, confType, err)
}

func KeyNotFoundError(key string) error {
	return fmt.Errorf("key %s not found in configuration", key)
}
