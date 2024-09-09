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

type UnfollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	FollowModel crud.FollowModel
}

func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
	return &UnfollowLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		FollowModel: crud.NewFollowModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// Unfollow 取消关注用户
func (l *UnfollowLogic) Unfollow(in *user_server.UnfollowRequest) (*user_server.UnfollowResponse, error) {
	// Step 1: 查询关注信息
	follow, err := l.FollowModel.FindOneByFollowerIdFolloweeId(l.ctx, in.FollowerId, in.FolloweeId)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("查询失败，followerId: %v; followeeId: %v", in.FollowerId, in.FolloweeId),
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

	return &user_server.UnfollowResponse{
		FolloweeId: follow.FolloweeId,
	}, nil
}
