package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok_e-commence/common"
	"tiktok_e-commence/common/model/model"
)

// @Summary 分页查询商品api
// @Tags 商品服务
// @Accept json
// @Produce json
// @Param user body model.ListProductsReq true "查询的商品和分页信息"
// @Success 200 {object} common.Response "查询成功"
// @Router /product/list [get]
func ListProductsHandler(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.ListProductsReq
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
		client := model.NewProductCatalogServiceClient(conn)
		log.Printf("Request: %+v", req)
		resp, err := client.ListProducts(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}

// @Summary 根据 id 查询商品api
// @Tags 商品服务
// @Accept json
// @Produce json
// @Param user body model.GetProductReq true "查询的商品 id"
// @Success 200 {object} common.Response "查询成功"
// @Router /product [get]
func GetProductHandler(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.GetProductReq
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
		client := model.NewProductCatalogServiceClient(conn)
		log.Printf("Request: %+v", req)
		resp, err := client.GetProduct(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}
