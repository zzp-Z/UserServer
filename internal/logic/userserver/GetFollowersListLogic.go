package userserverlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowersListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowersListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowersListLogic {
	return &GetFollowersListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取被关注列表
func (l *GetFollowersListLogic) GetFollowersList(in *user_server.GetFollowersListRequest) (*user_server.GetFollowersListResponse, error) {
	// todo: add your logic here and delete this line

	return &user_server.GetFollowersListResponse{}, nil
}
