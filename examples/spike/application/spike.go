package application

import (
	"cachemode/examples/spike/models"
	"cachemode/examples/spike/models/response"
	"cachemode/pkg/interfaces"
)

var (
	dbClient    interfaces.DB
	cacheClient interfaces.Cache
)

func init() {
	var err error
	dbClient, cacheClient, err = models.InitData()
	if err != nil {
		panic(err)
	}
	return
}

// 分页获取指定秒杀场景下的商品列表
func GetSpikeCommodityList(spikeID int, page, pageSize int) (commodities []response.CommodityResp, err error) {
	return
}

// 获取指定秒杀场景下的单个秒杀商品详情
func GetSpikeCommodityDetail(spikeID int, skuID string) (commodityDetail response.CommodityDetail, err error) {
	return
}

// 指定秒杀场景下的下单操作
// commodityMap key-value: skuID - count
func Spike(spikeID, userID int64, commodityMap map[int64]int) (err error) {
	return
}
