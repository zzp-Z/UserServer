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

type GetUserNonSensitiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	UserModel crud.UserModel
	Tools     *logic.Tools
}

func NewGetUserNonSensitiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserNonSensitiveLogic {
	return &GetUserNonSensitiveLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		UserModel: crud.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf),
		Tools:     &logic.Tools{},
	}
}

// GetUserNonSensitive 获取用户信息（不包含敏感信息）
func (l *GetUserNonSensitiveLogic) GetUserNonSensitive(in *UserService.GetUserInfoByIdRequest) (*UserService.UserNonSensitiveInfoResponse, error) {
	// Step 1: 查询用户信息
	user, err := l.UserModel.FindOne(l.ctx, in.UserId)
	fmt.Println(user)
	fmt.Println(err)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("userId: %d", in.UserId),
			Error:     err,
			ErrorCode: "GUNS4022",
		})
		return nil, err
	}

	return &UserService.UserNonSensitiveInfoResponse{
		UserId:   user.Id,
		Username: user.Username,
		Bio:      user.Bio,
		Quotes:   user.Quotes,
		MoodId:   uint64(user.MoodId),
	}, nil
}
