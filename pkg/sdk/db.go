package sdk

import (
	"cachemode/pkg/constant"
	"cachemode/pkg/db"
	"cachemode/pkg/interfaces"
)

func NewDBClient(dbType constant.DbType, addr string) interfaces.DB {
	return db.NewClientDB(dbType, addr)
}
