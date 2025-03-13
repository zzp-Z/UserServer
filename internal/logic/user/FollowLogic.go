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

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	FollowModel crud.FollowModel
	Tools       *logic.Tools
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		FollowModel: crud.NewFollowModel(svcCtx.SqlConn, svcCtx.CacheConf),
		Tools:       &logic.Tools{},
	}
}

// Follow 关注用户
func (l *FollowLogic) Follow(in *UserService.FollowRequest) (*UserService.FollowResponse, error) {
	// Step 1: 不能是自己
	if in.FollowingId == in.FollowerId {
		err := fmt.Errorf("不能关注自己")
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("FolloweeId: %v, FollowerId: %v", in.FollowingId, in.FollowerId),
			Error:     err,
			ErrorCode: "FL0922",
		})
		return nil, err
	}
	// Step 2: 不能重复关注
	_, err := l.FollowModel.FindOneByFollowerIdFollowingId(l.ctx, in.FollowerId, in.FollowingId)
	if err == nil {
		err := fmt.Errorf("不能重复关注")
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("FolloweeId: %v, FollowerId: %v", in.FollowingId, in.FollowerId),
			Error:     err,
			ErrorCode: "FL0921",
		})
		return nil, err
	}
	// Step 3: 插入数据
	_, err = l.FollowModel.Insert(l.ctx, &crud.Follow{
		FollowingId: in.FollowingId,
		FollowerId:  in.FollowerId,
	})
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("FolloweeId: %v, FollowerId: %v", in.FollowingId, in.FollowerId),
			Error:     err,
			ErrorCode: "FL0923",
		})
		return nil, err
	}

	return &UserService.FollowResponse{
		FollowingId: in.FollowingId,
	}, nil
}
