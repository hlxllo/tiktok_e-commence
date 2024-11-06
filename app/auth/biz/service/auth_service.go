package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tiktok_e-commence/app/auth/biz/model"
	"tiktok_e-commence/app/auth/biz/utils"
	"tiktok_e-commence/common"
)

// 实现服务端接口
type AuthServer struct {
	model.UnimplementedAuthServiceServer
}

// 实现 DeliverTokenByRPC
func (s *AuthServer) DeliverTokenByRPC(c context.Context, req *model.DeliverTokenReq) (*model.DeliveryResp, error) {
	token, err := utils.GenerateJWT(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, common.ErrGenerateJWTFailed)
	}
	return &model.DeliveryResp{Token: token}, nil
}
