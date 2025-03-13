package rolelogic

import (
	"context"
	"github.com/zzp-Z/UserServer/db/crud"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RoleModel crud.RoleModel
}

func NewCheckUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserRoleLogic {
	return &CheckUserRoleLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		RoleModel: crud.NewRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// CheckUserRole 检查用户角色
func (l *CheckUserRoleLogic) CheckUserRole(in *UserService.CheckUserRoleRequest) (*UserService.CheckUserRoleResponse, error) {
	// todo: add your logic here and delete this line

	return &UserService.CheckUserRoleResponse{}, nil
}
