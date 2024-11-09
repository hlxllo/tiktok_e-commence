package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok_e-commence/common"
	"tiktok_e-commence/common/model/model"
)

// @Summary 注册用户api
// @Tags 用户服务
// @Accept json
// @Produce json
// @Param user body model.RegisterReq true "新增的用户信息"
// @Success 200 {object} common.Response "注册成功"
// @Router /user/register [post]
func RegisterUserHandler(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.RegisterReq
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
		client := model.NewUserServiceClient(conn)
		req.Password = common.SHAEncoding(req.Password)
		req.ConfirmPassword = common.SHAEncoding(req.ConfirmPassword)
		log.Printf("Request: %+v", req)
		resp, err := client.Register(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, resp)
	}
}

// @Summary 登录用户api
// @Tags 用户服务
// @Accept json
// @Produce json
// @Param user body model.LoginReq true "登录的用户信息"
// @Success 200 {object} common.Response "登录成功"
// @Router /user/login [post]
func LoginUserHandler(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.LoginReq
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
		client := model.NewUserServiceClient(conn)
		req.Password = common.SHAEncoding(req.Password)
		log.Printf("Request: %+v", req)
		resp, err := client.Login(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		instances, err := common.SelectHealthyInstance("auth-server")
		if err != nil {
			common.HandleResponse(c, http.StatusInternalServerError, common.ErrDiscoverServiceFailed, nil)
			return
		}
		conn2, err := common.CreateGRPCConn(instances.Ip, int(instances.Port))
		defer conn.Close()
		authClient := model.NewAuthServiceClient(conn2)
		delResp, err := authClient.DeliverTokenByRPC(c, &model.DeliverTokenReq{UserId: resp.UserId})
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, delResp)
	}
}
