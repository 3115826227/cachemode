package sdk

import (
	"cachemode/pkg/constant"
	"testing"
)

func TestNewDBClient(t *testing.T) {
	NewDBClient(constant.MysqlDB, "")
}
