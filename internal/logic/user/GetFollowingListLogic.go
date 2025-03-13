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

type GetFollowingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	FollowModel crud.FollowModel
	Tools       *logic.Tools
}

func NewGetFollowingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingListLogic {
	return &GetFollowingListLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		FollowModel: crud.NewFollowModel(svcCtx.SqlConn, svcCtx.CacheConf),
		Tools:       &logic.Tools{},
	}
}

// GetFollowingList 获取关注列表
func (l *GetFollowingListLogic) GetFollowingList(in *UserService.GetFollowingListRequest) (*UserService.GetFollowingListResponse, error) {
	// Step 1: 获取关注列表
	followList, err := l.FollowModel.FindByFollowerId(l.ctx, in.UserId)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("UserId: %v", in.UserId),
			Error:     err,
			ErrorCode: "GDL821",
		})
		return nil, err
	}
	// Step 2: 创建切片，存放被关注列表的UserId
	userIds := make([]uint64, len(followList))
	for i, follow := range followList {
		userIds[i] = follow.FollowingId
	}

	return &UserService.GetFollowingListResponse{
		UserIds: userIds,
	}, nil
}
