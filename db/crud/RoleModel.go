package crud

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RoleModel = (*customRoleModel)(nil)

type (
	// RoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleModel.
	RoleModel interface {
		roleModel
		FindAll(ctx context.Context) ([]*Role, error)
	}

	customRoleModel struct {
		*defaultRoleModel
	}
)

// NewRoleModel returns a model for the database table.
func NewRoleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RoleModel {
	return &customRoleModel{
		defaultRoleModel: newRoleModel(conn, c, opts...),
	}
}

func (m *customRoleModel) FindAll(ctx context.Context) (r []*Role, err error) {
	query := fmt.Sprintf("select * from %s WHERE deleted_at IS NULL", m.table)
	err = m.QueryRowsNoCacheCtx(ctx, &r, query)
	if err != nil {
		return nil, err
	}
	return
}
