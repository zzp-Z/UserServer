package permissionlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPermissionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPermissionsLogic {
	return &GetUserPermissionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserPermissions 获取用户权限
func (l *GetUserPermissionsLogic) GetUserPermissions(in *UserService.GetUserPermissionsRequest) (*UserService.GetUserPermissionsResponse, error) {
	// todo: add your logic here and delete this line

	return &UserService.GetUserPermissionsResponse{}, nil
}
