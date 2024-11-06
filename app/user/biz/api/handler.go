package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
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
			st, ok := status.FromError(err)
			if ok {
				switch st.Code() {
				case codes.InvalidArgument:
					common.HandleResponse(c, http.StatusBadRequest, st.Message(), nil)
				case codes.AlreadyExists:
					common.HandleResponse(c, http.StatusBadRequest, st.Message(), nil)
				default:
					common.HandleResponse(c, http.StatusInternalServerError, st.Message(), nil)
				}
			} else {
				common.HandleResponse(c, http.StatusInternalServerError, err.Error(), nil)
			}
			return
		}
		common.HandleResponse(c, http.StatusOK, "success", resp)
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
			st, ok := status.FromError(err)
			if ok {
				switch st.Code() {
				case codes.NotFound:
					common.HandleResponse(c, http.StatusBadRequest, st.Message(), nil)
				default:
					common.HandleResponse(c, http.StatusInternalServerError, st.Message(), nil)
				}
			} else {
				common.HandleResponse(c, http.StatusInternalServerError, err.Error(), nil)
			}
			return
		}
		// 获取当前请求的上下文
		ctx := c.Request.Context()
		// 在上下文中设置值
		ctx = context.WithValue(ctx, "email", req.Email)
		ctx = context.WithValue(ctx, "password", req.Password)
		// TODO 这行之后删掉，去调用auth
		common.HandleResponse(c, http.StatusOK, "success", resp)
	}
}
