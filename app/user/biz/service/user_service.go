package service

import (
	"context"
	"fmt"
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
		// 参数错误
		return nil, status.Errorf(codes.InvalidArgument, common.ErrPasswordMismatch)
	}
	user := &model.User{Email: req.Email, Password: req.Password}
	userId, err := model.CreateUser(user)
	if err != nil {
		// 重复添加
		return nil, status.Errorf(codes.AlreadyExists, common.ErrUserExists)
	}
	return &model.RegisterResp{UserId: int32(userId)}, nil
}

// 实现 Login
func (s *UserServer) Login(c context.Context, req *model.LoginReq) (*model.LoginResp, error) {
	user := &model.User{Email: req.Email, Password: req.Password}
	userId, err := model.SelectUser(user)
	fmt.Println(err)
	if err != nil {
		// 没有用户
		return nil, status.Errorf(codes.NotFound, common.ErrLoginFailed)
	}
	return &model.LoginResp{UserId: int32(userId)}, nil
}
