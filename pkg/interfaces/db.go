package interfaces

// 关系型数据库操作接口定义
type DB interface {
	// 连接数据库操作
	Conn() error
	// 初始化建表操作
	InitTable(dos ...DataObject) error
	// 添加数据
	CreateObject(do DataObject) error
	// 判断数据是否存在
	ExistObject(query map[string]interface{}, do DataObject) (bool, error)
	// 删除数据
	DeleteObject(do DataObject) error
	// 更新数据
	UpdateObject(do DataObject) error
	// 查询单个数据
	GetObject(do DataObject) error
	// 查询列表数据
	GetObjects(query map[string]interface{}, do DataObject, likeParam *LikeQueryParam,
		order *OrderQueryParam, offset, limit int) ([]interface{}, int64, error)
}

// 查询数据单元接口定义
type DataObject interface {
	TableName() string
	Get() interface{}
}

// 模糊查询参数
type LikeQueryParam struct {
	Field  string
	Value  string
	Prefix bool
	Suffix bool
}

// 排序查询参数
type OrderQueryParam struct {
	Key  string
	Desc bool
}
