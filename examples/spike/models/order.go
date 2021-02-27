package models

// 订单
type Order struct {
	ID        int64 `gorm:"pk" json:"id"` // 订单编号
	Total     int   `json:"total"`        // 总商品数
	Price     int64 `json:"price"`        // 总价格
	UserID    int64 `json:"user_id"`      // 用户id
	PayStatus int   `json:"pay_status"`   // 用户支付状态
	Status    int   `json:"status"`       // 订单状态
	Time      `gorm:"embedded"`
}

func (table *Order) TableName() string {
	return "spike_order"
}

func (table *Order) Get() interface{} {
	return *table
}

// 订单与商品购买关联关系
type OrderCommodityRelation struct {
	OrderID        int64 `json:"order_id"`         // 订单id
	CommoditySkuID int64 `json:"commodity_sku_id"` // 商品sku_id
	Price          int64 `json:"price"`            // 购买的商品单价
	Count          int   `json:"count"`            // 购买的商品数
}

func (table *OrderCommodityRelation) TableName() string {
	return "spike_order_commodity_rel"
}

func (table *OrderCommodityRelation) Get() interface{} {
	return *table
}
