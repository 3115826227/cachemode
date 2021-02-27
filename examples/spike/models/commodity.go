package models

// 商品SPU(SPU 标准化产品单元)
type CommoditySPU struct {
	ID       int64  `gorm:"pk" json:"id"` // 商品SPU编号
	Name     string `json:"name"`         // 商品名称
	Title    string `json:"title"`        // 商品标题
	Saleable int    `json:"saleable"`     // 商品是否上架 0-未上架 1-上架
	Time     `gorm:"embedded"`
}

func (table *CommoditySPU) TableName() string {
	return "spike_commodity_spu"
}

func (table *CommoditySPU) Get() interface{} {
	return *table
}

// 商品SPU详情
type CommoditySPUDetail struct {
	SpuID       int64  `gorm:"pk" json:"spu_id"` // 商品SPU编号
	Description string `json:"description"`      // 商品描述
	Time        `gorm:"embedded"`
}

func (table *CommoditySPUDetail) TableName() string {
	return "spike_commodity_spu_detail"
}

func (table *CommoditySPUDetail) Get() interface{} {
	return *table
}

// 商品SKU(SKU 规格表)
type CommoditySKU struct {
	ID    int64  `gorm:"pk" json:"id"` // 商品SKU编号
	SpuID int64  `json:"spu_id"`       // 商品spu_id
	Title string `json:"title"`        // 商品标题
	Price int64  `json:"price"`        // 商品价格
	Time  `gorm:"embedded"`
}

func (table *CommoditySKU) TableName() string {
	return "spike_commodity_sku"
}

func (table *CommoditySKU) Get() interface{} {
	return *table
}

// 商品库存
type CommodityStock struct {
	SkuID int64 `json:"sku_id"` // 商品sku_id
	Stock int64 `json:"stock"`  // 商品库存数
	Time  `gorm:"embedded"`
}

func (table *CommodityStock) TableName() string {
	return "spike_commodity_stock"
}

func (table *CommodityStock) Get() interface{} {
	return *table
}

// 秒杀记录表
type CommoditySpike struct {
	ID         int64  `json:"id"`          // 秒杀编号
	SkuID      int64  `json:"sku_id"`      // 商品sku_id
	SpikeTitle string `json:"spkie_title"` // 秒杀标题
	SpikePrice int64  `json:"spike_price"` // 秒杀商品价格
	SpikeStock int64  `json:"spike_stock"` // 秒杀商品库存
	SpikeTotal int64  `json:"spike_total"` // 秒杀商品总数量
	Time       `gorm:"embedded"`
}

func (table *CommoditySpike) TableName() string {
	return "spike_commodity_spike"
}

func (table *CommoditySpike) Get() interface{} {
	return *table
}
