package models

// 用户
type User struct {
	ID       int64  `gorm:"pk" json:"id"` // 用户编号
	Username string `json:"username"`     // 用户名
	Time     `gorm:"embedded"`
}

func (table *User) TableName() string {
	return "spike_user"
}

func (table *User) Get() interface{} {
	return *table
}
