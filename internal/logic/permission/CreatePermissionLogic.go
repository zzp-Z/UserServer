package permissionlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePermissionLogic {
	return &CreatePermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreatePermission 创建权限
func (l *CreatePermissionLogic) CreatePermission(in *UserService.CreatePermissionRequest) (*UserService.CreatePermissionResponse, error) {
	// todo: add your logic here and delete this line

	return &UserService.CreatePermissionResponse{}, nil
}
