package permissionlogic

import (
	"context"
	"github.com/zzp-Z/UserServer/db/crud"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignRolePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PermissionModel     crud.PermissionModel
	RolePermissionModel crud.RolePermissionModel
}

func NewAssignRolePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignRolePermissionLogic {
	return &AssignRolePermissionLogic{
		ctx:                 ctx,
		svcCtx:              svcCtx,
		Logger:              logx.WithContext(ctx),
		PermissionModel:     crud.NewPermissionModel(svcCtx.SqlConn, svcCtx.CacheConf),
		RolePermissionModel: crud.NewRolePermissionModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// AssignRolePermission 角色权限
func (l *AssignRolePermissionLogic) AssignRolePermission(in *UserService.AssignRolePermissionRequest) (*UserService.AssignRolePermissionResponse, error) {
	// todo: add your logic here and delete this line

	return &UserService.AssignRolePermissionResponse{}, nil
}
