package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok_e-commence/app/order/biz/model"
	"tiktok_e-commence/common"
)

// @Summary 创建订单api
// @Tags 订单服务
// @Accept json
// @Produce json
// @Param user body model.PlaceOrderReq true "创建的订单信息"
// @Success 200 {object} common.Response "创建成功"
// @Router /order/create [post]
func PlaceOrderHandler(client model.OrderServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.PlaceOrderReq
		// 绑定参数
		err := c.ShouldBindJSON(&req)
		if err != nil {
			common.HandleResponse(c, http.StatusBadRequest, common.ErrInvalidParam, nil)
			return
		}
		// 打印请求参数以便调试
		log.Printf("Request: %+v", req)
		// 调用rpc分页查询
		resp, err := client.PlaceOrder(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}
