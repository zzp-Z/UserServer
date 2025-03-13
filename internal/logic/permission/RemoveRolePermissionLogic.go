package permissionlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveRolePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveRolePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveRolePermissionLogic {
	return &RemoveRolePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RemoveRolePermission 移除角色权限
func (l *RemoveRolePermissionLogic) RemoveRolePermission(in *UserService.RemoveRolePermissionRequest) (*UserService.RemoveRolePermissionResponse, error) {
	// todo: add your logic here and delete this line

	return &UserService.RemoveRolePermissionResponse{}, nil
}
