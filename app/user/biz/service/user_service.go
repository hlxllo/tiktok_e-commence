package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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
	queryUser := &model.User{Email: req.Email, Password: req.Password}
	user, err := model.SelectUser(queryUser)
	if err != nil {
		// 没有用户
		return nil, status.Errorf(codes.NotFound, common.ErrLoginFailed)
	}
	return &model.LoginResp{UserId: int32(user.ID)}, nil
}

// 实现 GetUserInfo  TODO 待测试
func (s *UserServer) GetUserInfo(c context.Context, req *model.GetUserInfoReq) (*model.GetUserInfoResp, error) {
	queryUser := &model.User{Model: gorm.Model{ID: uint(req.UserId)}, Email: req.Email}
	user, err := model.SelectUser(queryUser)
	if err != nil {
		// 没有用户
		return nil, status.Errorf(codes.NotFound, common.ErrUserNotFound)
	}
	return &model.GetUserInfoResp{UserId: int32(user.ID), Email: user.Email}, nil
}
