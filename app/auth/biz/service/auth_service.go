package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tiktok_e-commence/app/auth/biz/model"
)

// 实现服务端接口
type AuthServer struct {
	model.UnimplementedAuthServiceServer
}

// 实现 DeliverTokenByRPC
func (s *AuthServer) DeliverTokenByRPC(c context.Context, req *model.DeliverTokenReq) (*model.DeliveryResp, error) {
	email := c.Value("email")
	password := c.Value("password")
	return nil, status.Errorf(codes.Unimplemented, "method DeliverTokenByRPC not implemented")
}
