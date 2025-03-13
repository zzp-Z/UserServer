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

type UnfollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	FollowModel crud.FollowModel
	Tools       *logic.Tools
}

func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
	return &UnfollowLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		FollowModel: crud.NewFollowModel(svcCtx.SqlConn, svcCtx.CacheConf),
		Tools:       &logic.Tools{},
	}
}

// Unfollow 取消关注用户
func (l *UnfollowLogic) Unfollow(in *UserService.UnfollowRequest) (*UserService.UnfollowResponse, error) {
	// Step 1: 查询关注信息
	follow, err := l.FollowModel.FindOneByFollowerIdFollowingId(l.ctx, in.FollowerId, in.FollowingId)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("查询失败，followerId: %v; followeeId: %v", in.FollowerId, in.FollowingId),
			Error:     err,
			ErrorCode: "UF0193",
		})
		err = fmt.Errorf("没有关注信息")
		// 如果没有关注信息，则返回错误
		return nil, err
	}
	// Step 2: 删除关注信息
	err = l.FollowModel.Delete(l.ctx, follow.Id)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("删除失败，followId: %v", follow.Id),
			Error:     err,
			ErrorCode: "UF0194",
		})
		return nil, err
	}

	return &UserService.UnfollowResponse{
		FollowingId: follow.FollowingId,
	}, nil
}
