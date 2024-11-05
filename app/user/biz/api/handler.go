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

// 注册用户api
func RegisterUserHandler(client model.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.RegisterReq
		// 绑定参数
		err := c.ShouldBindJSON(&req)
		if err != nil {
			common.HandleResponse(c, http.StatusBadRequest, common.ErrInvalidParam, nil)
			return
		}
		// 打印请求参数以便调试
		log.Printf("Request: %+v", req)
		// 调用rpc注册用户
		resp, err := client.Register(context.Background(), &req)
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
