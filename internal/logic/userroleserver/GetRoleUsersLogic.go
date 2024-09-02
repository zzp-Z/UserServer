package userroleserverlogic

import (
	"context"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleUsersLogic {
	return &GetRoleUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取角色用户列表
func (l *GetRoleUsersLogic) GetRoleUsers(in *user_server.GetRoleUsersRequest) (*user_server.GetRoleUsersResponse, error) {
	// todo: add your logic here and delete this line

	return &user_server.GetRoleUsersResponse{}, nil
}
