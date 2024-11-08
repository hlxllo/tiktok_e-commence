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
// @Router /order [post]
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

// @Summary 批量查询订单api
// @Tags 订单服务
// @Accept json
// @Produce json
// @Param user body model.ListOrderReq true "查询的订单信息"
// @Success 200 {object} common.Response "查询成功"
// @Router /order [get]
func ListOrderHandler(client model.OrderServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.ListOrderReq
		// 绑定参数
		err := c.ShouldBindJSON(&req)
		if err != nil {
			common.HandleResponse(c, http.StatusBadRequest, common.ErrInvalidParam, nil)
			return
		}
		// 打印请求参数以便调试
		log.Printf("Request: %+v", req)
		// 调用rpc分页查询
		resp, err := client.ListOrder(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}

// @Summary 标记订单为已完成api
// @Tags 订单服务
// @Accept json
// @Produce json
// @Param user body model.MarkOrderPaidReq true "标记的订单信息"
// @Success 200 {object} common.Response "标记成功"
// @Router /order [delete]
func MarkOrderPaidHandler(client model.OrderServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.MarkOrderPaidReq
		// 绑定参数
		err := c.ShouldBindJSON(&req)
		if err != nil {
			common.HandleResponse(c, http.StatusBadRequest, common.ErrInvalidParam, nil)
			return
		}
		// 打印请求参数以便调试
		log.Printf("Request: %+v", req)
		// 调用rpc分页查询
		resp, err := client.MarkOrderPaid(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}
