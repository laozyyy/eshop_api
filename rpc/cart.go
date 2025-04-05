package rpc

import (
	"eshop_api/kitex_gen/eshop/cart"
	"eshop_api/kitex_gen/eshop/cart/cartservice"
	"eshop_api/log"
	"eshop_api/model/req"
	"eshop_api/model/resp"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
)

var cartService cartservice.Client

func init() {
	var err error
	cartService, err = cartservice.NewClient("hello", client.WithHostPorts("117.72.72.114:20002"))
	//cartService, err = cartservice.NewClient("hello", client.WithHostPorts("localhost:8888"))
	if err != nil {
		log.Errorf("error: %v", err)
	}
}

func AddItem(ctx *gin.Context, req req.AddItemRequestDTO) (*resp.AddItemRespDTO, error) {
	r := &cart.AddItemRequest{
		SkuId:    req.SkuId,
		Quantity: req.Quantity,
		Uid:      req.Uid,
	}
	item, err := cartService.AddItem(ctx, r)
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}
	ret := &resp.AddItemRespDTO{
		Code: int(item.Code),
	}
	if item.ErrStr != nil {
		ret.Info = *item.ErrStr
	} else {
		ret.Info = "success"
	}
	return ret, nil
}

func GetList(ctx *gin.Context, request cart.PageRequest) (*cart.PageResponse, error) {
	r := &cart.PageRequest{
		PageSize: request.PageSize,
		PageNum:  request.PageNum,
		Uid:      request.Uid,
	}
	list, err := cartService.GetList(ctx, r)
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}
	return list, nil
}

func UpdateItem(ctx *gin.Context, req req.UpdateItemRequestDTO) (*cart.UpdateResponse, error) {
	q := req.Quantity
	s := req.Selected
	r := &cart.UpdateRequest{
		SkuId:    req.SkuId,
		Quantity: &q,
		Selected: &s,
		Uid:      req.Uid,
	}
	res, err := cartService.UpdateItem(ctx, r)
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}
	return res, nil
}

func DeleteItem(ctx *gin.Context, req req.DeleteItemRequestDTO) (*resp.BaseRespDTO, error) {
	r := &cart.DeleteRequest{
		Skus: req.Skus,
		Uid:  req.Uid,
	}
	res, err := cartService.DeleteItem(ctx, r)
	if err != nil {
		log.Errorf("error: %v", err)
		return nil, err
	}
	return &resp.BaseRespDTO{
		Code: int(res.Code),
		Info: safeGetErrStr(res.ErrStr),
	}, nil
}

// 辅助函数处理可选错误信息
func safeGetErrStr(errStr *string) string {
	if errStr != nil {
		return *errStr
	}
	return "success"
}
