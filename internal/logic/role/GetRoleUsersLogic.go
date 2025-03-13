package rolelogic

import (
	"context"
	"github.com/zzp-Z/UserServer/db/crud"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RoleModel crud.RoleModel
}

func NewGetRoleUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleUsersLogic {
	return &GetRoleUsersLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		RoleModel: crud.NewRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// GetRoleUsers 获取角色用户列表
func (l *GetRoleUsersLogic) GetRoleUsers(in *UserService.GetRoleUsersRequest) (*UserService.GetRoleUsersResponse, error) {
	// todo: add your logic here and delete this line

	return &UserService.GetRoleUsersResponse{}, nil
}
