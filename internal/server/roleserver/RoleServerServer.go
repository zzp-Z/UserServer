// Code generated by goctl. DO NOT EDIT.
// Source: UserServer.proto

package server

import (
	"context"

	"github.com/zzp-Z/UserServer/internal/logic/roleserver"
	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"
)

type RoleServerServer struct {
	svcCtx *svc.ServiceContext
	user_server.UnimplementedRoleServerServer
}

func NewRoleServerServer(svcCtx *svc.ServiceContext) *RoleServerServer {
	return &RoleServerServer{
		svcCtx: svcCtx,
	}
}

// 获取角色信息
func (s *RoleServerServer) GetRole(ctx context.Context, in *user_server.GetRoleRequest) (*user_server.GetRoleResponse, error) {
	l := roleserverlogic.NewGetRoleLogic(ctx, s.svcCtx)
	return l.GetRole(in)
}

// 获取角色列表
func (s *RoleServerServer) GetRoleList(ctx context.Context, in *user_server.GetRoleListRequest) (*user_server.GetRoleListResponse, error) {
	l := roleserverlogic.NewGetRoleListLogic(ctx, s.svcCtx)
	return l.GetRoleList(in)
}

// 更新角色
func (s *RoleServerServer) UpdateRole(ctx context.Context, in *user_server.UpdateRoleRequest) (*user_server.UpdateRoleResponse, error) {
	l := roleserverlogic.NewUpdateRoleLogic(ctx, s.svcCtx)
	return l.UpdateRole(in)
}

// 删除角色
func (s *RoleServerServer) DeleteRole(ctx context.Context, in *user_server.DeleteRoleRequest) (*user_server.DeleteRoleResponse, error) {
	l := roleserverlogic.NewDeleteRoleLogic(ctx, s.svcCtx)
	return l.DeleteRole(in)
}

// 创建角色
func (s *RoleServerServer) CreateRole(ctx context.Context, in *user_server.CreateRoleRequest) (*user_server.CreateRoleResponse, error) {
	l := roleserverlogic.NewCreateRoleLogic(ctx, s.svcCtx)
	return l.CreateRole(in)
}
