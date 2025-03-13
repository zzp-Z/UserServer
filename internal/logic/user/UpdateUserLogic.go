package userlogic

import (
	"context"
	"fmt"
	"github.com/zzp-Z/UserServer/db/crud"
	"github.com/zzp-Z/UserServer/internal/logic"
	"github.com/zzp-Z/UserServer/log"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	UserModel crud.UserModel
	Tools     *logic.Tools
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		UserModel: crud.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf),
		Tools:     &logic.Tools{},
	}
}

// UpdateUser 更新用户信息
func (l *UpdateUserLogic) UpdateUser(in *UserService.UpdateUserRequest) (*UserService.UpdateUserResponse, error) {
	// Step 1: 查找用户
	user, err := l.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("userId: %d", in.UserId),
			Error:     err,
			ErrorCode: "UU3921",
		})
		return nil, err
	}
	// Step 2: 替换用户信息
	if in.Username != "" {
		user.Username = in.Username
	}
	if in.Bio != "" {
		user.Bio = in.Bio
	}
	// Step 3: 更新用户信息
	err = l.UserModel.Update(l.ctx, user)

	return &UserService.UpdateUserResponse{
		UserId: user.Id,
	}, nil
}
