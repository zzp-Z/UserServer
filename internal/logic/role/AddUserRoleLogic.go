package rolelogic

import (
	"context"
	"github.com/zzp-Z/UserServer/db/crud"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	UserRoleModel crud.UserRoleModel
}

func NewAddUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserRoleLogic {
	return &AddUserRoleLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		UserRoleModel: crud.NewUserRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// AddUserRole 添加用户角色
func (l *AddUserRoleLogic) AddUserRole(in *UserService.AddUserRoleRequest) (*UserService.AddUserRoleResponse, error) {
	// todo: add your logic here and delete this line

	return &UserService.AddUserRoleResponse{}, nil
}
