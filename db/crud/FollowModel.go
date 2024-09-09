package crud

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowModel = (*customFollowModel)(nil)

type (
	// FollowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowModel.
	FollowModel interface {
		followModel
		FindByFolloweeId(ctx context.Context, followeeId uint64) ([]*Follow, error)
		FindByFollowerId(ctx context.Context, followerId uint64) ([]*Follow, error)
	}

	customFollowModel struct {
		*defaultFollowModel
	}
)

// NewFollowModel returns a model for the database table.
func NewFollowModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FollowModel {
	return &customFollowModel{
		defaultFollowModel: newFollowModel(conn, c, opts...),
	}
}

// FindByFolloweeId 查询关注列表
func (m *customFollowModel) FindByFolloweeId(ctx context.Context, followeeId uint64) ([]*Follow, error) {
	query := fmt.Sprintf("select %s from %s where `followee_id` = ?", followRows, m.table)
	var follows []*Follow
	err := m.QueryRowsNoCacheCtx(ctx, &follows, query, followeeId)
	if err != nil {
		return nil, err
	}
	return follows, nil
}

// FindByFollowerId 根据followerId查询
func (m *customFollowModel) FindByFollowerId(ctx context.Context, followerId uint64) ([]*Follow, error) {
	query := fmt.Sprintf("select %s from %s where `follower_id` = ?", followRows, m.table)
	var follows []*Follow
	err := m.QueryRowsNoCacheCtx(ctx, &follows, query, followerId)
	if err != nil {
		return nil, err
	}
	return follows, nil
}
