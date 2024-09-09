package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zzp-Z/UserServer/internal/config"
	"os"
)

func init() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	// 如果 PodIp 字段不为空，则设置环境变量 POD_IP 为该值。
	if c.PodIp != "" {
		fmt.Println("上报etcd地址修改为: ", c.PodIp)
		_ = os.Setenv("POD_IP", c.PodIp)
	}
}
