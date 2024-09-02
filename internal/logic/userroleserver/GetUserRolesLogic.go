package userroleserverlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRolesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRolesLogic {
	return &GetUserRolesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户角色列表
func (l *GetUserRolesLogic) GetUserRoles(in *user_server.GetUserRolesRequest) (*user_server.GetUserRolesResponse, error) {
	// todo: add your logic here and delete this line

	return &user_server.GetUserRolesResponse{}, nil
}
