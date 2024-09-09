package followserverlogic

import (
	"context"
	"fmt"
	"github.com/zzp-Z/UserServer/db/crud"
	"github.com/zzp-Z/UserServer/log"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowersListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	FollowModel crud.FollowModel
}

func NewGetFollowersListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowersListLogic {
	return &GetFollowersListLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		FollowModel: crud.NewFollowModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// GetFollowersList 获取被关注列表
func (l *GetFollowersListLogic) GetFollowersList(in *user_server.GetFollowersListRequest) (*user_server.GetFollowersListResponse, error) {
	// Step 1: 获取被关注列表
	followList, err := l.FollowModel.FindByFolloweeId(l.ctx, in.UserId)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("UserId: %v", in.UserId),
			Error:     err,
			ErrorCode: "GDL831",
		})
		return nil, err
	}
	// Step 2: 创建切片，存放被关注列表的UserId
	userIds := make([]uint64, len(followList))
	for i, follow := range followList {
		userIds[i] = follow.FollowerId
	}
	return &user_server.GetFollowersListResponse{
		UserIds: userIds,
	}, nil
}
