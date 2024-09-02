package userserverlogic

import (
	"context"
	"database/sql"
	"github.com/zzp-Z/UserServer/db/crud"
	"github.com/zzp-Z/UserServer/log"
	"strconv"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	UserModel crud.UserModel
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		UserModel: crud.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// UpdateUser 更新用户信息
func (l *UpdateUserLogic) UpdateUser(in *user_server.UpdateUserRequest) (*user_server.UpdateUserResponse, error) {
	// Step 1: 查找用户
	user, err := l.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   strconv.FormatUint(in.UserId, 10),
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
		user.Bio = sql.NullString{
			String: in.Bio,
			Valid:  true,
		}
	}
	// Step 3: 更新用户信息
	err = l.UserModel.Update(l.ctx, user)

	return &user_server.UpdateUserResponse{
		UserId: user.Id,
	}, nil
}
