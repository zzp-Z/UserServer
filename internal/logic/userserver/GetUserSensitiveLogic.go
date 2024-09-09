package userserverlogic

import (
	"context"
	"fmt"
	"github.com/zzp-Z/UserServer/db/crud"
	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/log"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSensitiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	UserModel crud.UserModel
}

func NewGetUserSensitiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSensitiveLogic {
	return &GetUserSensitiveLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		UserModel: crud.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// GetUserSensitive 获取用户信息（包含敏感信息）
func (l *GetUserSensitiveLogic) GetUserSensitive(in *user_server.GetUserInfoByIdRequest) (*user_server.UserSensitiveInfoResponse, error) {
	// Step 1: 查询用户信息
	user, err := l.UserModel.FindOne(l.ctx, uint64(in.UserId))
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Error:     err,
			ErrorCode: "GUS711",
			Message:   fmt.Sprintf("userId: %d", in.UserId),
		})
		return nil, err
	}

	return &user_server.UserSensitiveInfoResponse{
		UserId:   user.Id,
		Username: user.Username,
		Bio:      user.Bio.String,
		Quotes:   user.Quotes.String,
		MoodId:   uint64(user.MoodId.Int64),
		Email:    user.Email,
	}, nil
}
