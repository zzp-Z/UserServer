package rolelogic

import (
	"context"
	"github.com/zzp-Z/UserServer/db/crud"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RoleModel crud.RoleModel
}

func NewDeleteUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserRoleLogic {
	return &DeleteUserRoleLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		RoleModel: crud.NewRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// DeleteUserRole 删除用户角色
func (l *DeleteUserRoleLogic) DeleteUserRole(in *UserService.DeleteUserRoleRequest) (*UserService.DeleteUserRoleResponse, error) {
	// todo: add your logic here and delete this line

	return &UserService.DeleteUserRoleResponse{}, nil
}
