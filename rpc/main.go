package rpc

import (
	"context"
	"eshop_api/kitex_gen/eshop/home"
	"eshop_api/kitex_gen/eshop/home/goodsservice"
	"eshop_api/log"
	"eshop_api/model"
	"github.com/cloudwego/kitex/client"
)

var goodsClient goodsservice.Client

func init() {
	var err error
	//goodsClient, err = goodsservice.NewClient("hello", client.WithHostPorts("localhost:8888"))
	goodsClient, err = goodsservice.NewClient("hello", client.WithHostPorts("117.72.72.114:20001"))
	if err != nil {
		log.Errorf("error: %v", err)
	}
}

func GetOneSku(ctx context.Context, sku string) (*model.GoodsSku, error) {
	resp, err := goodsClient.GetOneSku(ctx, sku)
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}
	if resp.Sku == nil {
		return nil, nil
	}

	return &model.GoodsSku{
		Sku:       resp.Sku.Sku,
		GoodsID:   resp.Sku.GoodsId,
		TagID:     resp.Sku.TagId,
		Name:      resp.Sku.Name,
		Price:     float64(resp.Sku.Price),
		Spec:      resp.Sku.Spec,
		ShowPic:   resp.Sku.ShowPic[0],
		DetailPic: resp.Sku.DetailPic[0],
	}, nil
}

func MGetSku(ctx context.Context, TagId string, pageSize int, pageNum int) ([]*model.GoodsSku, error) {
	request := home.MGetSkuRequest{
		PageSize: int32(pageSize),
		PageNum:  int32(pageNum),
		TagId:    TagId,
	}
	resp, err := goodsClient.MGetSku(ctx, &request)
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}
	if resp.Sku == nil {
		return nil, nil
	}

	result := make([]*model.GoodsSku, 0, len(resp.Sku))
	for _, sku := range resp.Sku {
		result = append(result, &model.GoodsSku{
			Sku:       sku.Sku,
			GoodsID:   sku.GoodsId,
			TagID:     sku.TagId,
			Name:      sku.Name,
			Price:     float64(sku.Price),
			Spec:      sku.Spec,
			ShowPic:   sku.ShowPic[0],
			DetailPic: sku.DetailPic[0],
		})
	}
	return result, nil
}

func GetRandomSku(ctx context.Context, pageSize int, pageNum int) ([]*model.GoodsSku, error) {
	request := home.PageRequest{
		PageSize: int32(pageSize),
		PageNum:  int32(pageNum),
	}
	resp, err := goodsClient.GetRandomSku(ctx, &request)
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}
	if resp.Sku == nil {
		return nil, nil
	}

	result := make([]*model.GoodsSku, 0, len(resp.Sku))
	for _, sku := range resp.Sku {
		result = append(result, &model.GoodsSku{
			Sku:       sku.Sku,
			GoodsID:   sku.GoodsId,
			TagID:     sku.TagId,
			Name:      sku.Name,
			Price:     float64(sku.Price),
			Spec:      sku.Spec,
			ShowPic:   sku.ShowPic[0],
			DetailPic: sku.DetailPic[0],
		})
	}
	return result, nil
}

func SearchGoods(ctx context.Context, keyword string, pageSize int, pageNum int) ([]*model.GoodsSku, error) {
	request := home.SearchRequest{
		Keyword:  keyword,
		PageSize: int32(pageSize),
		PageNum:  int32(pageNum),
	}
	resp, err := goodsClient.SearchGoods(ctx, &request)
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}

	result := make([]*model.GoodsSku, 0, len(resp.Sku))
	for _, sku := range resp.Sku {
		result = append(result, &model.GoodsSku{
			Sku:       sku.Sku,
			GoodsID:   sku.GoodsId,
			TagID:     sku.TagId,
			Name:      sku.Name,
			Price:     float64(sku.Price),
			Spec:      sku.Spec,
			ShowPic:   safeFirstElement(sku.ShowPic),
			DetailPic: safeFirstElement(sku.DetailPic),
		})
	}
	return result, nil
}

// 辅助函数处理空数组
func safeFirstElement(arr []string) string {
	if len(arr) > 0 {
		return arr[0]
	}
	return ""
}

func SubmitSeckillOrder(ctx context.Context, userID string, sku string, activityID string) (*home.SeckillResponse, error) {
	req := &home.SeckillRequest{
		UserId:     userID,
		Sku:        sku,
		ActivityId: activityID,
	}
	return goodsClient.SubmitSeckillOrder(ctx, req)
}

func DrawLottery(ctx context.Context, userID string, activityID string) (*home.LotteryResponse, error) {
	req := &home.LotteryRequest{
		UserId:     userID,
		ActivityId: activityID,
	}
	return goodsClient.DrawLottery(ctx, req)
}
