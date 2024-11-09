package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok_e-commence/common"
	"tiktok_e-commence/common/model/model"
)

// @Summary 新增购物车api
// @Tags 购物车服务
// @Accept json
// @Produce json
// @Param user body model.AddItemReq true "新增的购物车信息"
// @Success 200 {object} common.Response "新增成功"
// @Router /cart [post]
func AddItemHandler(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.AddItemReq
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
		client := model.NewCartServiceClient(conn)
		log.Printf("Request: %+v", req)
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
// @Param user body model.GetCartReq true "查询的购物车信息"
// @Success 200 {object} common.Response "查询成功"
// @Router /cart/get [post]
func GetCartHandler(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.GetCartReq
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
		client := model.NewCartServiceClient(conn)
		log.Printf("Request: %+v", req)
		resp, err := client.GetCart(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}

// @Summary 删除购物车api
// @Tags 购物车服务
// @Accept json
// @Produce json
// @Param user body model.EmptyCartReq true "删除的购物车信息"
// @Success 200 {object} common.Response "删除成功"
// @Router /cart [delete]
func EmptyCartHandler(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.EmptyCartReq
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
		client := model.NewCartServiceClient(conn)
		log.Printf("Request: %+v", req)
		resp, err := client.EmptyCart(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}
