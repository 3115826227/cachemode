package models

import (
	"testing"
)

func TestInitData(t *testing.T) {
	_, cacheClient, err := InitData()
	if err != nil {
		panic(err)
	}
	defer cacheClient.Close()
	return
}
