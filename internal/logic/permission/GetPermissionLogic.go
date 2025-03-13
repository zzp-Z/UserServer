package permissionlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionLogic {
	return &GetPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetPermission 获取权限信息
func (l *GetPermissionLogic) GetPermission(in *UserService.GetPermissionRequest) (*UserService.GetPermissionResponse, error) {
	// todo: add your logic here and delete this line

	return &UserService.GetPermissionResponse{}, nil
}
