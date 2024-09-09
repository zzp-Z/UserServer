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

type GetUserNonSensitiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	UserModel crud.UserModel
}

func NewGetUserNonSensitiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserNonSensitiveLogic {
	return &GetUserNonSensitiveLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		UserModel: crud.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// GetUserNonSensitive 获取用户信息（不包含敏感信息）
func (l *GetUserNonSensitiveLogic) GetUserNonSensitive(in *user_server.GetUserInfoByIdRequest) (*user_server.UserNonSensitiveInfoResponse, error) {
	// Step 1: 查询用户信息
	user, err := l.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("userId: %d", in.UserId),
			Error:     err,
			ErrorCode: "GUNS4022",
		})
		return nil, err
	}

	return &user_server.UserNonSensitiveInfoResponse{
		UserId:   user.Id,
		Username: user.Username,
		Bio:      user.Bio.String,
		Quotes:   user.Quotes.String,
		MoodId:   uint64(user.MoodId.Int64),
	}, nil
}
