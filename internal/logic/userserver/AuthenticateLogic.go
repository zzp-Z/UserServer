package userserverlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthenticateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticateLogic {
	return &AuthenticateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证用户密码（用途包括不限于“登录”，修改密码前的验证，修改信息前的验证等）
func (l *AuthenticateLogic) Authenticate(in *user_server.AuthenticateRequest) (*user_server.AuthenticateResponse, error) {
	// todo: add your logic here and delete this line

	return &user_server.AuthenticateResponse{}, nil
}
