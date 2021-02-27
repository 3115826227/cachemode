package mysql

import (
	"cachemode/pkg/constant"
	"cachemode/pkg/interfaces"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ClientMysql struct {
	addr string
	db   *gorm.DB
}

func NewClientMysql(addr string) interfaces.DB {
	return &ClientMysql{
		addr: addr,
	}
}

func (cm *ClientMysql) Conn() (err error) {
	cm.db, err = gorm.Open(mysql.Open(cm.addr), &gorm.Config{})
	return
}

func (cm *ClientMysql) InitTable(dos ...interfaces.DataObject) (err error) {
	var tables = make([]interface{}, 0)
	for _, do := range dos {
		tables = append(tables, do)
	}
	return cm.db.AutoMigrate(tables...)
}

// 添加
func (cm *ClientMysql) CreateObject(object interfaces.DataObject) (err error) {
	return cm.db.Create(object).Error
}

// 删除
func (cm *ClientMysql) DeleteObject(object interfaces.DataObject) (err error) {
	return cm.db.Delete(object).Error
}

// 获取结果
func (cm *ClientMysql) GetObject(object interfaces.DataObject) (err error) {
	return cm.db.Table(object.TableName()).First(object).Error
}

// 更新数据
func (cm *ClientMysql) UpdateObject(object interfaces.DataObject) (err error) {
	return cm.db.Table(object.TableName()).Save(object).Error
}

// 判断是否存在
func (cm *ClientMysql) ExistObject(query map[string]interface{}, do interfaces.DataObject) (exist bool, err error) {
	var count int64
	template := cm.db.Table(do.TableName())
	for key, value := range query {
		template = template.Where(fmt.Sprintf("%v = ?", key), value)
	}
	err = template.First(do).Count(&count).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return
	}
	exist = true
	return
}

// 条件查询
func (cm *ClientMysql) GetObjects(query map[string]interface{}, do interfaces.DataObject, likeParam *interfaces.LikeQueryParam, order *interfaces.OrderQueryParam, offset, limit int) (list []interface{}, count int64, err error) {
	template := cm.db.Table(do.TableName())
	for key, value := range query {
		template = template.Where(fmt.Sprintf("%v = ?", key), value)
	}
	if likeParam != nil {
		var value = likeParam.Value
		if likeParam.Prefix {
			value = "%" + value
		}
		if likeParam.Suffix {
			value = value + "%"
		}
		template = template.Where(fmt.Sprintf("%v LIKE ?", likeParam.Field), value)
	}
	if err = template.Count(&count).Error; err != nil {
		return
	}
	if order != nil {
		orderKey := order.Key
		if order.Desc {
			orderKey = fmt.Sprintf("%v %v", orderKey, constant.DESCORDER)
		} else {
			orderKey = fmt.Sprintf("%v %v", orderKey, constant.ASCORDER)
		}
		template = template.Order(orderKey)
	}
	rows, err := template.Limit(limit).Offset(offset).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		if err = cm.db.ScanRows(rows, do); err != nil {
			return
		}
		list = append(list, do.Get())
	}
	return
}
