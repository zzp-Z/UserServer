package userserverlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSensitiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserSensitiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSensitiveLogic {
	return &GetUserSensitiveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息（包含敏感信息）
func (l *GetUserSensitiveLogic) GetUserSensitive(in *user_server.GetUserInfoByIdRequest) (*user_server.UserSensitiveInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user_server.UserSensitiveInfoResponse{}, nil
}
