package confgetter

import (
	"bitbucket.org/HeilaSystems/configurations/config"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"time"
)

type confgetterWrapper struct {
	instance ConfGetter
}
func (v *confgetterWrapper) Get(key string) config.Value {
	return &valueWrapper{
		key:               key,
		confgetterWrapper: v,
	}
}

func (v *confgetterWrapper) Implementation() interface{} {
	return v.instance
}

const keyNotFound = "keyNotFound"

type valueWrapper struct {
	key string
	*confgetterWrapper
}

func (v *valueWrapper) IsSet() bool {
	_, ok := v.instance[v.key]
	return ok
}

func (v *valueWrapper) Raw() (interface{},error) {
	val, ok := v.instance[v.key]
	if !ok {
		return nil,fmt.Errorf(keyNotFound)
	}
	return val,nil
}

func (v *valueWrapper) Bool() (bool,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return false,fmt.Errorf(keyNotFound)
	}
	return cast.ToBoolE(val)
}

func (v *valueWrapper) Int() (int,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0,fmt.Errorf(keyNotFound)
	}
	return cast.ToIntE(val)
}

func (v *valueWrapper) Int32() (int32,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0,fmt.Errorf(keyNotFound)
	}
	return cast.ToInt32E(val)
}

func (v *valueWrapper) Int64() (int64,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0,fmt.Errorf(keyNotFound)
	}
	return cast.ToInt64E(val)
}

func (v *valueWrapper) Uint() (uint,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0,fmt.Errorf(keyNotFound)
	}
	return cast.ToUintE(val)
}

func (v *valueWrapper) Uint32() (uint32,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0,fmt.Errorf(keyNotFound)
	}
	return cast.ToUint32E(val)
}

func (v *valueWrapper) Uint64() (uint64,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0,fmt.Errorf(keyNotFound)
	}
	return cast.ToUint64E(val)
}

func (v *valueWrapper) Float64() (float64,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0,fmt.Errorf(keyNotFound)
	}
	return cast.ToFloat64E(val)
}

func (v *valueWrapper) Time() (time.Time,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return time.Time{} ,fmt.Errorf(keyNotFound)
	}
	return cast.ToTimeE(val)
}

func (v *valueWrapper) Duration() (time.Duration,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return 0 ,fmt.Errorf(keyNotFound)
	}
	return cast.ToDurationE(val)
}

func (v *valueWrapper) String() (string,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return "" ,fmt.Errorf(keyNotFound)
	}
	return cast.ToStringE(val)
}

func (v *valueWrapper) IntSlice() ([]int,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return nil ,fmt.Errorf(keyNotFound)
	}
	return cast.ToIntSliceE(val)
}

func (v *valueWrapper) StringSlice() ([]string,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return nil ,fmt.Errorf(keyNotFound)
	}
	return cast.ToStringSliceE(val)
}

func (v *valueWrapper) StringMap() (map[string]interface{},error) {
	val, ok := v.instance[v.key]
	if !ok {
		return nil ,fmt.Errorf(keyNotFound)
	}
	return cast.ToStringMapE(val)
}

func (v *valueWrapper) StringMapString() (map[string]string,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return nil ,fmt.Errorf(keyNotFound)
	}
	return cast.ToStringMapStringE(val)
}

func (v *valueWrapper) StringMapStringSlice() (map[string][]string,error) {
	val, ok := v.instance[v.key]
	if !ok {
		return nil ,fmt.Errorf(keyNotFound)
	}
	return cast.ToStringMapStringSliceE(val)
}

// Unmarshal uses default decoder options, if you need some special behavior than it's best to get cfg.Implementation() and use it from there
func (v *valueWrapper) Unmarshal(result interface{}) error {
	val, ok := v.instance[v.key]
	if !ok {
		return fmt.Errorf(keyNotFound)
	}
	if byteArray , err := json.Marshal(val);err != nil{
		return err
	}else {
		return json.Unmarshal(byteArray,&result)
	}
}
