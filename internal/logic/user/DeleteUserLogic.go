package userlogic

import (
	"context"
	"fmt"
	"github.com/zzp-Z/UserServer/db/crud"
	"github.com/zzp-Z/UserServer/internal/logic"
	"github.com/zzp-Z/UserServer/log"
	"time"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	UserModel crud.UserModel
	Tools     *logic.Tools
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		UserModel: crud.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf),
		Tools:     &logic.Tools{},
	}
}

// DeleteUser 删除用户
func (l *DeleteUserLogic) DeleteUser(in *UserService.DeleteUserRequest) (*UserService.DeleteUserResponse, error) {
	// Step 1: 查找用户
	user, err := l.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("UserId: %d", in.UserId),
			Error:     err,
			ErrorCode: "DU823",
		})
		return nil, err
	}
	// Step 2: 删除用户
	user.DeletedAt = time.Time{}

	err = l.UserModel.Update(l.ctx, user)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("UserId: %d", in.UserId),
			Error:     err,
			ErrorCode: "DU824",
		})
		return nil, err
	}

	return &UserService.DeleteUserResponse{
		UserId: user.Id,
	}, nil
}
