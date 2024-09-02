package userroleserverlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserRoleLogic {
	return &DeleteUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除用户角色
func (l *DeleteUserRoleLogic) DeleteUserRole(in *user_server.DeleteUserRoleRequest) (*user_server.DeleteUserRoleResponse, error) {
	// todo: add your logic here and delete this line

	return &user_server.DeleteUserRoleResponse{}, nil
}
