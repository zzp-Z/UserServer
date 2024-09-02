package userserverlogic

import (
	"context"
	"fmt"
	"github.com/zzp-Z/UserServer/db/crud"
	"github.com/zzp-Z/UserServer/internal/logic"
	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/log"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	UserModel crud.UserModel
	Tools     *logic.Tools
}

func NewAuthenticateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticateLogic {
	return &AuthenticateLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		UserModel: crud.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf),
		Tools:     &logic.Tools{},
	}
}

// Authenticate 验证用户密码（用途包括不限于“登录”，修改密码前的验证，修改信息前的验证等）
func (l *AuthenticateLogic) Authenticate(in *user_server.AuthenticateRequest) (*user_server.AuthenticateResponse, error) {
	// Step 1: 验证邮箱是否存在
	user, err := l.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   in.Email,
			Error:     err,
			ErrorCode: "AA900",
		})
		return nil, fmt.Errorf("邮箱不存在")
	}
	// Step 2: 验证密码是否正确
	isValid := l.Tools.CheckPassword(in.Password, user.HashPassword)
	if !isValid {
		err = fmt.Errorf("密码错误")
		log.Error(nil, log.ErrorContent{
			Message:   in.Email,
			Error:     err,
			ErrorCode: "AA901",
		})
		return nil, err
	}

	return &user_server.AuthenticateResponse{
		UserId: user.Id,
	}, nil
}
