package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok_e-commence/app/cart/biz/model"
	"tiktok_e-commence/common"
)

// @Summary 新增购物车api
// @Tags 购物车服务
// @Accept json
// @Produce json
// @Param user body model.AddItemReqCopy true "新增的购物车信息"
// @Success 200 {object} common.Response "查询成功"
// @Router /cart/create [post]
func AddItemHandler(client model.CartServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.AddItemReq
		// 绑定参数
		err := c.ShouldBindJSON(&req)
		if err != nil {
			common.HandleResponse(c, http.StatusBadRequest, common.ErrInvalidParam, nil)
			return
		}
		// 打印请求参数以便调试
		log.Printf("Request: %+v", req)
		// 调用rpc
		resp, err := client.AddItem(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}

// @Summary 查询购物车api
// @Tags 购物车服务
// @Accept json
// @Produce json
// @Param user body model.GetCartReqCopy true "查询的购物车信息"
// @Success 200 {object} common.Response "查询成功"
// @Router /cart/get [post]
func GetCartHandler(client model.CartServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.GetCartReq
		// 绑定参数
		err := c.ShouldBindJSON(&req)
		if err != nil {
			common.HandleResponse(c, http.StatusBadRequest, common.ErrInvalidParam, nil)
			return
		}
		// 打印请求参数以便调试
		log.Printf("Request: %+v", req)
		// 调用rpc
		resp, err := client.GetCart(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}
