package db

import (
	"cachemode/pkg/constant"
	"cachemode/pkg/db/mysql"
	"cachemode/pkg/interfaces"
)

type ClientDB struct {
	client interfaces.DB
}

func NewClientDB(dbType constant.DbType, addr string) *ClientDB {
	clientDB := new(ClientDB)
	switch dbType {
	case constant.MysqlDB:
		clientDB.client = mysql.NewClientMysql(addr)
	}
	return clientDB
}

func (cdb *ClientDB) Conn() error {
	return cdb.client.Conn()
}

func (cdb *ClientDB) InitTable(dos ...interfaces.DataObject) error {
	return cdb.client.InitTable(dos...)
}

// 添加数据
func (cdb *ClientDB) CreateObject(do interfaces.DataObject) error {
	return cdb.client.CreateObject(do)
}

// 判断数据是否存在
func (cdb *ClientDB) ExistObject(query map[string]interface{}, do interfaces.DataObject) (bool, error) {
	return cdb.client.ExistObject(query, do)
}

// 删除数据
func (cdb *ClientDB) DeleteObject(do interfaces.DataObject) error {
	return cdb.client.InitTable(do)
}

// 更新数据
func (cdb *ClientDB) UpdateObject(do interfaces.DataObject) error {
	return cdb.client.UpdateObject(do)
}

// 查询单个数据
func (cdb *ClientDB) GetObject(do interfaces.DataObject) error {
	return cdb.client.GetObject(do)
}

// 查询列表数据
func (cdb *ClientDB) GetObjects(query map[string]interface{}, do interfaces.DataObject, likeParam *interfaces.LikeQueryParam,
	order *interfaces.OrderQueryParam, offset, limit int) ([]interface{}, int64, error) {
	return cdb.client.GetObjects(query, do, likeParam, order, offset, limit)
}
