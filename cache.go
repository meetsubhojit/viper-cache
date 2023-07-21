package pkg

import (
	"time"

	"github.com/jellydator/ttlcache/v3"
)

var viperCache *ttlcache.Cache[string, viperValue]

const (
	defaultTTL = 5 * time.Second
)

type viperValue struct {
	value interface{}
}

func init() {
	viperCache = ttlcache.New[string, viperValue](
		ttlcache.WithTTL[string, viperValue](defaultTTL),
	)
}

func get(key string) (interface{}, bool) {
	v := viperCache.Get(key)
	if v != nil {
		return v.Value().value, true
	}
	return nil, false
}

func set(key string, value viperValue) {
	viperCache.Set(key, value, defaultTTL)
}
