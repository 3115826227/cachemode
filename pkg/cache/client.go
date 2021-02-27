package cache

import (
	"cachemode/pkg/cache/redis"
	"cachemode/pkg/constant"
	"cachemode/pkg/interfaces"
)

type ClientCache struct {
	cacheType constant.CacheType
	client    interfaces.Cache
}

func NewClientCache(cacheType constant.CacheType, addr string) *ClientCache {
	clientCache := new(ClientCache)
	clientCache.cacheType = cacheType
	switch cacheType {
	case constant.RedisCache:
		clientCache.client = redis.NewClientRedis(addr)
	}
	return clientCache
}
