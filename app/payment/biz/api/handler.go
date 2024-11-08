package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok_e-commence/app/payment/biz/model"
	"tiktok_e-commence/common"
)

// @Summary 创建支付记录api
// @Tags 支付服务
// @Accept json
// @Produce json
// @Param user body model.ChargeReqCopy true "创建的支付信息"
// @Success 200 {object} common.Response "创建成功"
// @Router /payment [post]
func ChargeHandler(client model.PaymentServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.ChargeReq
		// 绑定参数
		err := c.ShouldBindJSON(&req)
		if err != nil {
			common.HandleResponse(c, http.StatusBadRequest, common.ErrInvalidParam, nil)
			return
		}
		// 打印请求参数以便调试
		log.Printf("Request: %+v", req)
		// 调用rpc分页查询
		resp, err := client.Charge(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}
