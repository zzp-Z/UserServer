package userserverlogic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zzp-Z/UserServer/db/crud"
	"github.com/zzp-Z/UserServer/log"
	"time"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	UserModel crud.UserModel
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		UserModel: crud.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// DeleteUser 删除用户
func (l *DeleteUserLogic) DeleteUser(in *user_server.DeleteUserRequest) (*user_server.DeleteUserResponse, error) {
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
	user.DeletedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err = l.UserModel.Update(l.ctx, user)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("UserId: %d", in.UserId),
			Error:     err,
			ErrorCode: "DU824",
		})
		return nil, err
	}

	return &user_server.DeleteUserResponse{
		UserId: user.Id,
	}, nil
}
