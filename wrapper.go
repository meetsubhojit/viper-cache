package viper_cache

import (
	"time"

	"github.com/spf13/viper"
)

func getAny[V any](key string, defV V, viperGetFunc func(key string) V) (retV V) {
	item, found := get(key)
	if found {
		if ret, ok := item.(V); ok {
			return ret
		}
	} else {
		vItem := viperGetFunc(key)
		set(key, viperValue{value: vItem})
		return vItem
	}
	return defV
}

func Get(key string) interface{} {
	return getAny[any](key, nil, viper.Get)
}
func GetBool(key string) bool {
	return getAny[bool](key, false, viper.GetBool)
}
func GetDuration(key string) time.Duration {
	return getAny[time.Duration](key, 0, viper.GetDuration)
}
func GetFloat64(key string) float64 {
	return getAny[float64](key, 0, viper.GetFloat64)
}
func GetInt(key string) int {
	return getAny[int](key, 0, viper.GetInt)
}
func GetInt32(key string) int32 {
	return getAny[int32](key, 0, viper.GetInt32)
}
func GetInt64(key string) int64 {
	return getAny[int64](key, 0, viper.GetInt64)
}
func GetIntSlice(key string) []int {
	return getAny[[]int](key, nil, viper.GetIntSlice)
}
func GetString(key string) string {
	return getAny[string](key, "", viper.GetString)
}
func GetStringSlice(key string) []string {
	return getAny[[]string](key, nil, viper.GetStringSlice)
}
func GetStringMap(key string) map[string]interface{} {
	return getAny[map[string]interface{}](key, nil, viper.GetStringMap)
}
func GetStringMapString(key string) map[string]string {
	return getAny[map[string]string](key, nil, viper.GetStringMapString)
}
func GetTime(key string) time.Time {
	return getAny[time.Time](key, time.Time{}, viper.GetTime)
}
func GetUint(key string) uint {
	return getAny[uint](key, 0, viper.GetUint)
}
func GetUint16(key string) uint16 {
	return getAny[uint16](key, 0, viper.GetUint16)
}
func GetUint32(key string) uint32 {
	return getAny[uint32](key, 0, viper.GetUint32)
}
func GetUint64(key string) uint64 {
	return getAny[uint64](key, 0, viper.GetUint64)
}

/*
func viperTest() {
	viper.Get("")
	viper.GetBool()
	viper.GetDuration()
	viper.GetFloat64()
	viper.GetInt()
	viper.GetInt32()
	viper.GetInt64()
	// viper.GetIntSlice()
	viper.GetString()
	viper.GetStringMap()
	viper.GetStringMapString()

	viper.GetTime()
	viper.GetUint()
	viper.GetStringSlice()
	viper.GetIntSlice()
	viper.Get

}
*/
