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
func CheckoutHandler(client model.CheckoutServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.CheckoutReq
		// 绑定参数
		err := c.ShouldBindJSON(&req)
		if err != nil {
			common.HandleResponse(c, http.StatusBadRequest, common.ErrInvalidParam, nil)
			return
		}
		// 打印请求参数以便调试
		log.Printf("Request: %+v", req)
		// 调用rpc分页查询
		resp, err := client.Checkout(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}
