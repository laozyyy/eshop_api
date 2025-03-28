package handler

import (
	"eshop_api/kitex_gen/eshop/cart"
	"eshop_api/log"
	"eshop_api/model/req"
	"eshop_api/model/resp"
	"eshop_api/rpc"
	"github.com/gin-gonic/gin"
)

func HandleAddItem(ctx *gin.Context) {
	var req req.AddItemRequestDTO
	var response *resp.AddItemRespDTO
	if err := ctx.ShouldBind(&req); err != nil {
		log.Errorf("error: %v", err)
		return
	}

	// RPC 调用
	response, err := rpc.AddItem(ctx, req)
	if err != nil {
		log.Errorf("error: %v", err)
		return
	}
	ctx.JSON(200, response)
}

func HandleCartGetList(ctx *gin.Context) {
	var req cart.PageRequest
	if err := ctx.ShouldBind(&req); err != nil {
		// 错误处理...
		return
	}

	// RPC 调用
	items, err := rpc.GetList(ctx, req)
	if err != nil {
		log.Errorf("error: %v", err)
		return
	}

	response := &resp.CartListRespDTO{
		Code: 200,
		Info: "success",
		Data: items,
	}
	ctx.JSON(200, response)
}
