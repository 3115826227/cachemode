package mysql

import (
	"gorm.io/gorm"
)

type ClientMysql struct {
	db *gorm.DB
}
