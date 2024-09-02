package followserverlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingListLogic {
	return &GetFollowingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取关注列表
func (l *GetFollowingListLogic) GetFollowingList(in *user_server.GetFollowingListRequest) (*user_server.GetFollowingListResponse, error) {
	// todo: add your logic here and delete this line

	return &user_server.GetFollowingListResponse{}, nil
}
