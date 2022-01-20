package usecase

import (
	"logbook-api/internal/config"
	"logbook-api/internal/infrastructure"
	"logbook-api/pkg/cache"
	"strconv"

	"github.com/umatare5/atmos-go"
)

// Usecase struct
type Usecase struct {
	Config     *config.Config
	Repository *infrastructure.Repository
}

var inMemoryCache = cache.New(cache.DefaultExpirationTime, cache.DefaultCleanupInterval)

// GetDivelog returns the message
func (u *Usecase) GetDivelog(divelogID string) (atmos.GetDivelogResponse, error) {
	key := cache.DivelogEndpointCacheKey

	// Try to load cache
	data, cached := inMemoryCache.Get(key)
	if cached {
		return data.(atmos.GetDivelogResponse), nil
	}

	// Request divelog
	res, err := u.Repository.InvokeAtmosRepository().GetDivelog(divelogID)
	if err != nil {
		return *res, err
	}

	// Cache divelog
	inMemoryCache.Set(key, *res, cache.CommonExpirationTime)
	return *res, nil
}

// GetDivelogs returns the message
func (u *Usecase) GetDivelogs(limit *int, cursor *string) (atmos.GetDivelogsResponse, error) {
	key := u.buildDivelogsEndpointCacheKey(limit, cursor)

	// Try to load cache
	data, cached := inMemoryCache.Get(key)
	if cached {
		return data.(atmos.GetDivelogsResponse), nil
	}

	// Request divelogs
	res, err := u.Repository.InvokeAtmosRepository().GetDivelogs(limit, cursor)
	if err != nil {
		return *res, err
	}

	// Cache divelogs
	inMemoryCache.Set(key, *res, cache.CommonExpirationTime)
	return *res, nil
}

func (u *Usecase) buildDivelogsEndpointCacheKey(limit *int, cursor *string) string {
	if limit == nil && cursor == nil {
		return cache.DivelogsEndpointCacheKey
	}
	if cursor == nil {
		return cache.DivelogsEndpointCacheKey + strconv.Itoa(*limit)
	}
	if limit == nil {
		return cache.DivelogsEndpointCacheKey + *cursor
	}

	return cache.DivelogsEndpointCacheKey + strconv.Itoa(*limit) + *cursor
}
