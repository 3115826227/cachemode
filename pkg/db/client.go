package db

import "cachemode/pkg/interfaces"

type ClientDB struct {
	client interfaces.DB
}

func NewClientDB() *ClientDB {
	return &ClientDB{}
}
