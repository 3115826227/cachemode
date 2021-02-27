package constant

type DbType string

const (
	MysqlDB DbType = "mysql"
)

type OrderType string

const (
	// 顺序排序
	ASCORDER OrderType = "asc"
	// 倒序排序
	DESCORDER OrderType = "desc"
)
