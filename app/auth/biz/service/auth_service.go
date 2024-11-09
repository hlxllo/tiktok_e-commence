package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tiktok_e-commence/app/auth/biz/utils"
	"tiktok_e-commence/common"
	"tiktok_e-commence/common/model/model"
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

// 实现 VerifyTokenByRPC
func (s *AuthServer) VerifyTokenByRPC(c context.Context, req *model.VerifyTokenReq) (*model.VerifyResp, error) {
	resp := &model.VerifyResp{Res: false}
	// 解析jwt
	claims, err := utils.VerifyJWT(req.Token)
	if err != nil {
		return resp, status.Errorf(codes.Internal, common.ErrVerifyJWTFailed)
	}
	// 获取用户id
	id := claims["id"].(float64)
	instances, err := common.SelectHealthyInstance("user-server")
	if err != nil {
		return resp, status.Errorf(codes.Internal, common.ErrDiscoverServiceFailed)
	}
	conn, err := common.CreateGRPCConn(instances.Ip, int(instances.Port))
	defer conn.Close()
	userClient := model.NewUserServiceClient(conn)
	// 根据 id 查询用户信息
	_, err = userClient.GetUserInfo(c, &model.GetUserInfoReq{UserId: int32(id)})
	if err != nil {
		return resp, status.Errorf(codes.NotFound, common.ErrUserNotFound)
	}
	return &model.VerifyResp{Res: true}, nil
}
