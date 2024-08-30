package userserverlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserNonSensitiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserNonSensitiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserNonSensitiveLogic {
	return &GetUserNonSensitiveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息（不包含敏感信息）
func (l *GetUserNonSensitiveLogic) GetUserNonSensitive(in *user_server.GetUserInfoByIdRequest) (*user_server.UserNonSensitiveInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user_server.UserNonSensitiveInfoResponse{}, nil
}
