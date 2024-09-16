package roleserverlogic

import (
	"context"
	"github.com/zzp-Z/UserServer/db/crud"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RoleModel crud.RoleModel
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		RoleModel: crud.NewRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// GetRoleList 获取角色列表
func (l *GetRoleListLogic) GetRoleList() (*user_server.GetRoleListResponse, error) {
	// Step 1: 获取角色列表
	all, err := l.RoleModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	// Step 2: 转换为pb格式
	list := make([]*user_server.GetRoleResponse, len(all))
	for i, v := range all {
		list[i] = &user_server.GetRoleResponse{
			RoleId:   v.Id,
			RoleName: v.Name,
		}
	}
	// Step 3: 返回
	return &user_server.GetRoleListResponse{
		Roles: list,
	}, nil
}
