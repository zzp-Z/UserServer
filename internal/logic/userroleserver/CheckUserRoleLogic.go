package userroleserverlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserRoleLogic {
	return &CheckUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 检查用户角色
func (l *CheckUserRoleLogic) CheckUserRole(in *user_server.CheckUserRoleRequest) (*user_server.CheckUserRoleResponse, error) {
	// todo: add your logic here and delete this line

	return &user_server.CheckUserRoleResponse{}, nil
}
