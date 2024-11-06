package common

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

// 公共错误处理函数
func HandleError(c *gin.Context, err error) {
	st, ok := status.FromError(err)
	if ok {
		switch st.Code() {
		case codes.InvalidArgument:
			HandleResponse(c, http.StatusBadRequest, st.Message(), nil)
		case codes.AlreadyExists:
			HandleResponse(c, http.StatusBadRequest, st.Message(), nil)
		case codes.NotFound:
			HandleResponse(c, http.StatusBadRequest, st.Message(), nil)
		default:
			HandleResponse(c, http.StatusInternalServerError, st.Message(), nil)
		}
	} else {
		HandleResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
}
