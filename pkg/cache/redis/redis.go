package redis

import (
	"cachemode/pkg/interfaces"
	"github.com/gomodule/redigo/redis"
)

type ClientRedis struct {
	addr string
	conn redis.Conn
}

func NewClientRedis(addr string) *ClientRedis {
	return &ClientRedis{
		addr: addr,
	}
}

func (rc *ClientRedis) Conn() error {
	conn, err := redis.Dial("tcp", rc.addr)
	if err != nil {
		return err
	}
	rc.conn = conn
	return nil
}

func (rc *ClientRedis) Write(key interfaces.CacheKey, value []byte) (err error) {
	return
}

func (rc *ClientRedis) Read(key interfaces.CacheKey) (value []byte, err error) {
	return
}

func (rc *ClientRedis) Close() error {
	return rc.conn.Close()
}
