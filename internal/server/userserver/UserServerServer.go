// Code generated by goctl. DO NOT EDIT.
// Source: UserServer.proto

package server

import (
	"context"

	"github.com/zzp-Z/UserServer/internal/logic/userserver"
	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"
)

type UserServerServer struct {
	svcCtx *svc.ServiceContext
	user_server.UnimplementedUserServerServer
}

func NewUserServerServer(svcCtx *svc.ServiceContext) *UserServerServer {
	return &UserServerServer{
		svcCtx: svcCtx,
	}
}

// 注册新用户
func (s *UserServerServer) Register(ctx context.Context, in *user_server.RegisterRequest) (*user_server.RegisterResponse, error) {
	l := userserverlogic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

// 验证用户密码（用途包括不限于“登录”，修改密码前的验证，修改信息前的验证等）
func (s *UserServerServer) Authenticate(ctx context.Context, in *user_server.AuthenticateRequest) (*user_server.AuthenticateResponse, error) {
	l := userserverlogic.NewAuthenticateLogic(ctx, s.svcCtx)
	return l.Authenticate(in)
}

// 获取用户信息（包含敏感信息）
func (s *UserServerServer) GetUserSensitive(ctx context.Context, in *user_server.GetUserInfoByIdRequest) (*user_server.UserSensitiveInfoResponse, error) {
	l := userserverlogic.NewGetUserSensitiveLogic(ctx, s.svcCtx)
	return l.GetUserSensitive(in)
}

// 获取用户信息（不包含敏感信息）
func (s *UserServerServer) GetUserNonSensitive(ctx context.Context, in *user_server.GetUserInfoByIdRequest) (*user_server.UserNonSensitiveInfoResponse, error) {
	l := userserverlogic.NewGetUserNonSensitiveLogic(ctx, s.svcCtx)
	return l.GetUserNonSensitive(in)
}

// 更新用户信息
func (s *UserServerServer) UpdateUser(ctx context.Context, in *user_server.UpdateUserRequest) (*user_server.UpdateUserResponse, error) {
	l := userserverlogic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

// 修改密码
func (s *UserServerServer) ChangePassword(ctx context.Context, in *user_server.ChangePasswordRequest) (*user_server.ChangePasswordResponse, error) {
	l := userserverlogic.NewChangePasswordLogic(ctx, s.svcCtx)
	return l.ChangePassword(in)
}

// 删除用户
func (s *UserServerServer) DeleteUser(ctx context.Context, in *user_server.DeleteUserRequest) (*user_server.DeleteUserResponse, error) {
	l := userserverlogic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}
