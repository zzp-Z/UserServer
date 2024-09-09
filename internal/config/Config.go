package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PodIp      string
	DataSource string // 数据库连接字符串
	Cache      cache.CacheConf
}
