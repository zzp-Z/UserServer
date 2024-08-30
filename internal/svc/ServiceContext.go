package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zzp-Z/UserServer/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	SqlConn   sqlx.SqlConn
	CacheConf cache.CacheConf
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:    c,
		SqlConn:   conn,
		CacheConf: c.Cache, // 确保 CacheConf 被正确传递
	}
}
