package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tiktok_e-commence/app/user/biz/model"
	"tiktok_e-commence/common"
)

// 实现服务端接口
type UserServer struct {
	model.UnimplementedUserServiceServer
}

// 实现 Register
func (s *UserServer) Register(c context.Context, req *model.RegisterReq) (*model.RegisterResp, error) {
	if req.Password != req.ConfirmPassword {
		return nil, status.Errorf(codes.InvalidArgument, common.ErrPasswordMismatch)
	}
	user := &model.User{Email: req.Email, Password: req.Password}
	userId, err := model.CreateUser(user)
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, common.ErrUserExists)
	}
	return &model.RegisterResp{UserId: int32(userId)}, nil
}

// TODO 实现 Login
func (s *UserServer) Login(context.Context, *model.LoginReq) (*model.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
