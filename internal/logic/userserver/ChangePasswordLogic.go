package userserverlogic

import (
	"context"
	"github.com/zzp-Z/UserServer/db/crud"
	"github.com/zzp-Z/UserServer/internal/logic"
	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/log"
	"github.com/zzp-Z/UserServer/user_server"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	UserModel crud.UserModel
	Tools     *logic.Tools
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		UserModel: crud.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf),
		Tools:     &logic.Tools{},
	}
}

// ChangePassword 修改密码
func (l *ChangePasswordLogic) ChangePassword(in *user_server.ChangePasswordRequest) (*user_server.ChangePasswordResponse, error) {
	// Step 1: 查找用户
	user, err := l.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   strconv.FormatUint(in.UserId, 10),
			Error:     err,
			ErrorCode: "CP491",
		})
		return nil, err
	}
	// Step 2: 密码hash
	hashedPassword, err := l.Tools.HashPassword(in.NewPassword)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   in.NewPassword,
			Error:     err,
			ErrorCode: "CP492",
		})
		return nil, err
	}
	// Step 3: 更新密码
	user.HashPassword = hashedPassword
	err = l.UserModel.Update(l.ctx, user)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   hashedPassword,
			Error:     err,
			ErrorCode: "CP491",
		})
		return nil, err
	}

	return &user_server.ChangePasswordResponse{
		UserId: user.Id,
	}, nil
}
