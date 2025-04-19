package handler

import (
	"eshop_api/common/constant"
	"eshop_api/model/req"
	"eshop_api/model/resp"
	"eshop_api/rpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleSubmitSeckill(ctx *gin.Context) {
	var req req.SeckillReqDTO
	var res resp.SeckillRespDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res = resp.SeckillRespDTO{
			Code: constant.ParamError,
			Info: "参数错误: " + err.Error(),
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	result, err := rpc.SubmitSeckillOrder(ctx.Request.Context(), req.UserID, req.Sku, req.ActivityID)
	if err != nil {
		res = resp.SeckillRespDTO{
			Code: constant.ServerError,
			Info: "服务器内部错误",
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	res = resp.SeckillRespDTO{
		Code:    constant.Success,
		Info:    result.Info,
		OrderID: result.OrderId,
	}
	ctx.JSON(http.StatusOK, res)
}

func HandleDrawLottery(ctx *gin.Context) {
	var req req.LotteryReqDTO
	var res resp.LotteryRespDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res = resp.LotteryRespDTO{
			Code: constant.ParamError,
			Info: "参数错误: " + err.Error(),
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	result, err := rpc.DrawLottery(ctx.Request.Context(), req.UserID, req.ActivityID)
	if err != nil {
		res = resp.LotteryRespDTO{
			Code: constant.ServerError,
			Info: "服务器内部错误",
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	res = resp.LotteryRespDTO{
		Code:    constant.Success,
		HasWon:  result.HasWon,
		Sku:     result.Sku,
		Info:    result.Info,
		OrderID: result.OrderId,
	}
	ctx.JSON(http.StatusOK, res)
}
