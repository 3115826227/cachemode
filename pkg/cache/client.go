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

func (cc *ClientCache) Conn() error {
	return cc.client.Conn()
}

func (cc *ClientCache) Write(key interfaces.CacheKey, value []byte) (err error) {
	return cc.client.Write(key, value)
}

func (cc *ClientCache) Read(key interfaces.CacheKey) (value []byte, err error) {
	return cc.client.Read(key)
}

func (cc *ClientCache) Close() error {
	return cc.client.Close()
}
