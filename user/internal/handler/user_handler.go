package handler

import (
	"context"
	"user/internal/dao"
	"user/internal/service"
	"user/pkg/e"
)

type UserService struct {
	service.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user *dao.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.SUCCESS
	user, err = dao.CheckUserExists(req.UserName)
	if err != nil {
		resp.Code = e.ERROR
		return resp, err
	}
	resp.UserDetail = dao.BuildUser(user)
	return resp, nil
}

func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user *dao.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.SUCCESS
	user, err = dao.UserCreate(req)
	if err != nil {
		resp.Code = e.ERROR
		return resp, err
	}
	resp.UserDetail = dao.BuildUser(user)
	return resp, nil
}
