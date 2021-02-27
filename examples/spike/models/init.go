package models

import (
	"cachemode/pkg/constant"
	"cachemode/pkg/interfaces"
	"cachemode/pkg/sdk"
)

var (
	mysqlDbAddr    = "root:123456@tcp(127.0.0.1:13306)/cachemode?charset=utf8mb4&parseTime=True&loc=Local"
	redisCacheAddr = "localhost:16379"
)

func InitData() (dbClient interfaces.DB, cacheClient interfaces.Cache, err error) {
	dbClient = sdk.NewDBClient(constant.MysqlDB, mysqlDbAddr)
	if err = dbClient.Conn(); err != nil {
		return
	}
	if err = dbClient.InitTable(
		&CommoditySPU{},
		&CommoditySPUDetail{},
		&CommoditySKU{},
		&CommodityStock{},
		&CommoditySpike{},
		&User{},
		&Order{},
		&OrderCommodityRelation{},
	); err != nil {
		return
	}
	cacheClient = sdk.NewCacheClient(constant.RedisCache, redisCacheAddr)
	if err = cacheClient.Conn(); err != nil {
		return
	}
	return
}
