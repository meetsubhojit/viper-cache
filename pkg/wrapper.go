package pkg

import (
	"time"

	"github.com/spf13/viper"
)

func getAny[V any](key string, defV V) V {
	item, found := get(key)
	if found {
		if ret, ok := item.(V); ok {
			return ret
		}
	} else {
		vItem := viper.Get(key)
		if vItem != nil {
			if vItemT, ok := vItem.(V); ok {
				set(key, viperValue{value: vItemT})
				return vItemT
			}
		}
	}
	return defV
}

func Get(key string) interface{} {
	return getAny[any](key, nil)
}
func GetBool(key string) bool {
	return getAny[bool](key, false)
}
func GetDuration(key string) time.Duration {
	return getAny[time.Duration](key, 0)
}
func GetFloat64(key string) float64 {
	return getAny[float64](key, 0)
}
func GetInt(key string) int {
	return getAny[int](key, 0)
}
func GetInt32(key string) int32 {
	return getAny[int32](key, 0)
}
func GetInt64(key string) int64 {
	return getAny[int64](key, 0)
}
func GetIntSlice(key string) []int {
	return getAny[[]int](key, nil)
}
func GetString(key string) string {
	return getAny[string](key, "")
}
func GetStringSlice(key string) []string {
	return getAny[[]string](key, nil)
}
func GetStringMap(key string) map[string]interface{} {
	return getAny[map[string]interface{}](key, nil)
}
func GetStringMapString(key string) map[string]string {
	return getAny[map[string]string](key, nil)
}
func GetTime(key string) time.Time {
	return getAny[time.Time](key, time.Time{})
}
func GetUint(key string) uint {
	return getAny[uint](key, 0)
}
func GetUint16(key string) uint16 {
	return getAny[uint16](key, 0)
}
func GetUint32(key string) uint32 {
	return getAny[uint32](key, 0)
}
func GetUint64(key string) uint64 {
	return getAny[uint64](key, 0)
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
