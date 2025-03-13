package permissionlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolePermissionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRolePermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolePermissionsLogic {
	return &GetRolePermissionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetRolePermissions 获取角色权限
func (l *GetRolePermissionsLogic) GetRolePermissions(in *UserService.GetRolePermissionsRequest) (*UserService.GetRolePermissionsResponse, error) {
	// todo: add your logic here and delete this line

	return &UserService.GetRolePermissionsResponse{}, nil
}
