package rolelogic

import (
	"context"
	"github.com/zzp-Z/UserServer/db/crud"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRolesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RoleModel crud.RoleModel
}

func NewGetUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRolesLogic {
	return &GetUserRolesLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		RoleModel: crud.NewRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// GetUserRoles 获取用户角色列表
func (l *GetUserRolesLogic) GetUserRoles(in *UserService.GetUserRolesRequest) (*UserService.GetUserRolesResponse, error) {
	// todo: add your logic here and delete this line

	return &UserService.GetUserRolesResponse{}, nil
}
