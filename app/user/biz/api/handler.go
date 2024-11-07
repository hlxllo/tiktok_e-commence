package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	model2 "tiktok_e-commence/app/auth/biz/model"
	"tiktok_e-commence/app/user/biz/model"
	"tiktok_e-commence/common"
)

// @Summary 注册用户api
// @Tags 用户服务
// @Accept json
// @Produce json
// @Param user body model.RegisterReq true "新增的用户信息"
// @Success 200 {object} common.Response "注册成功"
// @Router /user/register [post]
func RegisterUserHandler(client model.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.RegisterReq
		// 绑定参数
		err := c.ShouldBindJSON(&req)
		if err != nil {
			common.HandleResponse(c, http.StatusBadRequest, common.ErrInvalidParam, nil)
			return
		}
		// 密码加密
		req.Password = common.SHAEncoding(req.Password)
		req.ConfirmPassword = common.SHAEncoding(req.ConfirmPassword)
		// 打印请求参数以便调试
		log.Printf("Request: %+v", req)
		// 调用rpc注册用户
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
func LoginUserHandler(client model.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.LoginReq
		// 绑定参数
		err := c.ShouldBindJSON(&req)
		if err != nil {
			common.HandleResponse(c, http.StatusBadRequest, common.ErrInvalidParam, nil)
			return
		}
		// 密码加密
		req.Password = common.SHAEncoding(req.Password)
		// 打印请求参数以便调试
		log.Printf("Request: %+v", req)
		// 调用rpc登录用户
		resp, err := client.Login(c, &req)
		if err != nil {
			common.HandleError(c, err)
			return
		}
		// 选择实例
		instances, err := common.SelectHealthyInstance("auth-server")
		if err != nil {
			common.HandleResponse(c, http.StatusInternalServerError, common.ErrDiscoverServiceFailed, nil)
			return
		}
		// 创建 grpc 客户端连接
		conn, err := common.CreateGRPCConn(instances.Ip, int(instances.Port))
		defer conn.Close()
		// 建立连接
		authClient := model2.NewAuthServiceClient(conn)
		delResp, err := authClient.DeliverTokenByRPC(c, &model2.DeliverTokenReq{UserId: resp.UserId})
		if err != nil {
			common.HandleError(c, err)
			return
		}
		common.HandleResponse(c, http.StatusOK, common.MsgSuccess, delResp)
	}
}
