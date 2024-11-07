package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok_e-commence/app/product/biz/model"
	"tiktok_e-commence/common"
)

// @Summary 分页查询商品api
// @Tags 商品服务
// @Accept json
// @Produce json
// @Param user body model.ListProductsReq true "查询的商品和分页信息"
// @Success 200 {object} common.Response "查询成功"
// @Router /product/list [post]
func ListProductsHandler(client model.ProductCatalogServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.ListProductsReq
		// 绑定参数
		err := c.ShouldBindJSON(&req)
		if err != nil {
			common.HandleResponse(c, http.StatusBadRequest, common.ErrInvalidParam, nil)
			return
		}
		// 打印请求参数以便调试
		log.Printf("Request: %+v", req)
		// 调用rpc分页查询
		resp, err := client.ListProducts(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}
