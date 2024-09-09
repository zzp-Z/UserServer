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

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	FollowModel crud.FollowModel
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		FollowModel: crud.NewFollowModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// Follow 关注用户
func (l *FollowLogic) Follow(in *user_server.FollowRequest) (*user_server.FollowResponse, error) {
	// Step 1: 不能是自己
	if in.FolloweeId == in.FollowerId {
		err := fmt.Errorf("不能关注自己")
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("FolloweeId: %v, FollowerId: %v", in.FolloweeId, in.FollowerId),
			Error:     err,
			ErrorCode: "FL0922",
		})
		return nil, err
	}
	// Step 2: 不能重复关注
	_, err := l.FollowModel.FindOneByFollowerIdFolloweeId(l.ctx, in.FollowerId, in.FolloweeId)
	if err == nil {
		err := fmt.Errorf("不能重复关注")
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("FolloweeId: %v, FollowerId: %v", in.FolloweeId, in.FollowerId),
			Error:     err,
			ErrorCode: "FL0921",
		})
		return nil, err
	}
	// Step 3: 插入数据
	_, err = l.FollowModel.Insert(l.ctx, &crud.Follow{
		FolloweeId: in.FolloweeId,
		FollowerId: in.FollowerId,
	})
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("FolloweeId: %v, FollowerId: %v", in.FolloweeId, in.FollowerId),
			Error:     err,
			ErrorCode: "FL0923",
		})
		return nil, err
	}

	return &user_server.FollowResponse{
		FolloweeId: in.FolloweeId,
	}, nil
}
