package permissionlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePermissionLogic {
	return &DeletePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeletePermission 删除权限
func (l *DeletePermissionLogic) DeletePermission(in *UserService.DeletePermissionRequest) (*UserService.DeletePermissionResponse, error) {
	// todo: add your logic here and delete this line

	return &UserService.DeletePermissionResponse{}, nil
}
