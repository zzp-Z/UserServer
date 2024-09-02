package userroleserverlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserRoleLogic {
	return &AddUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加用户角色
func (l *AddUserRoleLogic) AddUserRole(in *user_server.AddUserRoleRequest) (*user_server.AddUserRoleResponse, error) {
	// todo: add your logic here and delete this line

	return &user_server.AddUserRoleResponse{}, nil
}
