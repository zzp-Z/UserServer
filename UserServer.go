package main

import (
	"flag"
	"fmt"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/config"
	permissionServer "github.com/zzp-Z/UserServer/internal/server/permission"
	roleServer "github.com/zzp-Z/UserServer/internal/server/role"
	userServer "github.com/zzp-Z/UserServer/internal/server/user"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/UserServer.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		UserService.RegisterUserServer(grpcServer, userServer.NewUserServer(ctx))
		UserService.RegisterRoleServer(grpcServer, roleServer.NewRoleServer(ctx))
		UserService.RegisterPermissionServer(grpcServer, permissionServer.NewPermissionServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
