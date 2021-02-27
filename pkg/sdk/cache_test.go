package sdk

import (
	"cachemode/pkg/constant"
	"testing"
)

var (
	redisAddr = "localhost:6379"
)

func TestNewCacheClient(t *testing.T) {
	client := NewCacheClient(constant.RedisCache, redisAddr)
	if err := client.Conn(); err != nil {
		panic(err)
	}
}
