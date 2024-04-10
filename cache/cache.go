package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var C *cache.Cache

func init() {
	C = cache.New(5*time.Minute, 10*time.Minute)
}

// Set adds an item to the cache, replacing any existing item.
func Set(key string, value interface{}) {
	C.Set(key, value, cache.DefaultExpiration)
}

// Get retrieves an item from the cache.
func Get(key string) (interface{}, bool) {
	return C.Get(key)
}

// Delete removes an item from the cache.
func Delete(key string) {
	C.Delete(key)
}
