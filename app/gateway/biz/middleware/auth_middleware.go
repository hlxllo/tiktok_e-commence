package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok_e-commence/common"
	"tiktok_e-commence/common/model/model"
)

func AuthMiddleware(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		token := c.GetHeader("Authorization")
		if token == "" {
			common.HandleResponse(c, http.StatusUnauthorized, common.ErrVerifyJWTFailed, nil)
			c.Abort()
			return
		}
		// 调用登录校验服务
		conn, err := common.CallService(c, serviceName)
		if err != nil {
			c.Abort()
			return
		}
		defer conn.Close()
		client := model.NewAuthServiceClient(conn)
		resp, err := client.VerifyTokenByRPC(context.Background(), &model.VerifyTokenReq{Token: token})
		if err != nil || !resp.Res {
			common.HandleResponse(c, http.StatusUnauthorized, common.ErrVerifyJWTFailed, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
