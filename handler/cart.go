package handler

import (
	"eshop_api/common/constant"
	"eshop_api/kitex_gen/eshop/cart"
	"eshop_api/log"
	"eshop_api/model/req"
	"eshop_api/model/resp"
	"eshop_api/rpc"
	"github.com/gin-gonic/gin"
)

func HandleCartAddItem(ctx *gin.Context) {
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
	rpcResp, err := rpc.GetList(ctx, req)
	if err != nil {
		log.Errorf("error: %v", err)
		return
	}

	response := &resp.CartListRespDTO{
		Code:  200,
		Info:  "success",
		Data:  rpcResp.Items,
		Price: rpcResp.Price,
	}
	ctx.JSON(200, response)
}

func HandleCartUpdate(ctx *gin.Context) {
	var req req.UpdateItemRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, resp.BaseRespDTO{
			Code: constant.ParamError,
			Info: "参数错误: " + err.Error(),
		})
		return
	}

	response, err := rpc.UpdateItem(ctx, req)
	if err != nil {
		ctx.JSON(500, resp.BaseRespDTO{
			Code: constant.ServerError,
			Info: "服务调用失败",
		})
		return
	}
	ctx.JSON(200, resp.UpdateRespDTO{
		BaseRespDTO: resp.BaseRespDTO{
			Code: int(response.Code),
			Info: "success",
		},
		Price: response.Price,
	})
}

func HandleCartDelete(ctx *gin.Context) {
	var req req.DeleteItemRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, resp.BaseRespDTO{
			Code: constant.ParamError,
			Info: "参数错误: " + err.Error(),
		})
		return
	}

	response, err := rpc.DeleteItem(ctx, req)
	if err != nil {
		ctx.JSON(500, resp.BaseRespDTO{
			Code: constant.ServerError,
			Info: "服务调用失败",
		})
		return
	}
	ctx.JSON(200, response)
}
