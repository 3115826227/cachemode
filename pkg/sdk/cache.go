package sdk

import (
	"cachemode/pkg/cache"
	"cachemode/pkg/constant"
	"cachemode/pkg/interfaces"
)

func NewCacheClient(cacheType constant.CacheType, addr string) interfaces.Cache {
	return cache.NewClientCache(cacheType, addr)
}
