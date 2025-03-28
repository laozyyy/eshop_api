package handler

import (
	"eshop_api/common/constant"
	"eshop_api/model"
	"eshop_api/model/req"
	"eshop_api/model/resp"
	"eshop_api/rpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetOneSku(ctx *gin.Context) {
	var res resp.GetSkuRespDTO

	sku := ctx.Param("sku")
	if sku == "" {
		res = resp.GetSkuRespDTO{
			Code: constant.ParamError,
			Info: "sku不能为空",
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	goods, err := rpc.GetOneSku(ctx.Request.Context(), sku)
	if err != nil {
		res = resp.GetSkuRespDTO{
			Code: constant.ServerError,
			Info: "服务器内部错误",
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	if goods == nil {
		res = resp.GetSkuRespDTO{
			Code: constant.NotFound,
			Info: "商品不存在",
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	res = resp.GetSkuRespDTO{
		Code: constant.Success,
		Info: "success",
		Data: goods,
	}
	ctx.JSON(http.StatusOK, res)
}

func HandleMGetSku(ctx *gin.Context) {
	var (
		param req.MGetSkuReqDTO
		res   resp.PageRespDTO
	)

	if err := ctx.ShouldBindJSON(&param); err != nil {
		res = resp.PageRespDTO{
			Code: constant.ParamError,
			Info: "参数错误: " + err.Error(),
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	goods, err := rpc.MGetSku(ctx.Request.Context(), param.TagID, param.PageSize, param.PageNum)
	if err != nil {
		res = resp.PageRespDTO{
			Code: constant.ServerError,
			Info: "服务器内部错误",
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	if goods == nil {
		res = resp.PageRespDTO{
			Code: constant.Success,
			Info: "无数据",
			Data: []*model.GoodsSku{},
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	res = resp.PageRespDTO{
		Code: constant.Success,
		Info: "success",
		Data: goods,
	}
	ctx.JSON(http.StatusOK, res)
}

func HandleRandom(ctx *gin.Context) {
	var (
		param req.PageReqDTO
		res   resp.PageRespDTO
	)

	if err := ctx.ShouldBindJSON(&param); err != nil {
		res = resp.PageRespDTO{
			Code: constant.ParamError,
			Info: "参数错误: " + err.Error(),
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	goods, err := rpc.GetRandomSku(ctx.Request.Context(), param.PageSize, param.PageNum)
	if err != nil {
		res = resp.PageRespDTO{
			Code: constant.ServerError,
			Info: "服务器内部错误",
		}
		ctx.JSON(http.StatusOK, res)
		return

	}

	if goods == nil {
		res = resp.PageRespDTO{
			Code: constant.Success,
			Info: "无数据",
			Data: []*model.GoodsSku{},
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	res = resp.PageRespDTO{
		Code: constant.Success,
		Info: "success",
		Data: goods,
	}
	ctx.JSON(http.StatusOK, res)

}

func HandleSearchGoods(ctx *gin.Context) {
	var param req.SearchReqDTO
	var res resp.PageRespDTO

	if err := ctx.ShouldBindJSON(&param); err != nil {
		res = resp.PageRespDTO{
			Code: constant.ParamError,
			Info: "参数错误: " + err.Error(),
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	goods, err := rpc.SearchGoods(ctx.Request.Context(), param.Keyword, param.PageSize, param.PageNum)
	if err != nil {
		res = resp.PageRespDTO{
			Code: constant.ServerError,
			Info: "服务器内部错误",
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	res = resp.PageRespDTO{
		Code: constant.Success,
		Info: "success",
		Data: goods,
	}
	ctx.JSON(http.StatusOK, res)
}
