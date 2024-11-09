package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok_e-commence/common"
	"tiktok_e-commence/common/model/model"
)

// @Summary 创建支付记录api
// @Tags 支付服务
// @Accept json
// @Produce json
// @Param user body model.ChargeReq true "创建的支付信息"
// @Success 200 {object} common.Response "创建成功"
// @Router /payment [post]
func ChargeHandler(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.ChargeReq
		err := c.ShouldBindJSON(&req)
		if err != nil {
			common.HandleResponse(c, http.StatusBadRequest, common.ErrInvalidParam, nil)
			return
		}
		conn, err := common.CallService(c, serviceName)
		if err != nil {
			return
		}
		defer conn.Close()
		client := model.NewPaymentServiceClient(conn)
		log.Printf("Request: %+v", req)
		resp, err := client.Charge(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}
