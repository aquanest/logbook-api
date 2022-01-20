package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Default constants
const (
	DefaultExpirationTime    = 5 * time.Minute
	DefaultCleanupInterval   = 10 * time.Minute
	CommonExpirationTime     = 6 * time.Hour
	DivelogEndpointCacheKey  = "divelog"
	DivelogsEndpointCacheKey = "divelogs"
)

// New reserve memory for cache
func New(defaultExpiration, cleanupInterval time.Duration) *cache.Cache {
	return cache.New(defaultExpiration, cleanupInterval)
}
