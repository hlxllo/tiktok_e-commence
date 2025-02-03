package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok_e-commence/common"
	"tiktok_e-commence/common/model/model"
)

// @Summary 创建结算记录api
// @Tags 结算服务
// @Accept json
// @Produce json
// @Param user body model.CheckoutReq true "创建的结算信息"
// @Success 200 {object} common.Response "创建成功"
// @Router /checkout [post]
func CheckoutHandler(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.CheckoutReq
		// TODO 还有问题
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
		client := model.NewCheckoutServiceClient(conn)
		log.Printf("Request: %+v", req)
		resp, err := client.Checkout(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}
