package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var C *cache.Cache

func init() {
	C = cache.New(5*time.Minute, 10*time.Minute)
}

func Set(key string, value interface{}) {
	C.Set(key, value, cache.DefaultExpiration)
}

func Get(key string) (interface{}, bool) {
	return C.Get(key)
}
