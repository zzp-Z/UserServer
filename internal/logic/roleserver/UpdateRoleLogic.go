package roleserverlogic

import (
	"context"
	"fmt"
	"github.com/zzp-Z/UserServer/db/crud"
	"github.com/zzp-Z/UserServer/log"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RoleModel crud.RoleModel
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		RoleModel: crud.NewRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// UpdateRole 更新角色
func (l *UpdateRoleLogic) UpdateRole(in *user_server.UpdateRoleRequest) (*user_server.UpdateRoleResponse, error) {
	// Step 1: 查询角色
	role, err := l.RoleModel.FindOne(l.ctx, in.RoleId)
	// 当查询不到或者已经被删除
	if err != nil || role.DeletedAt.Valid == true {
		if err == nil { // 没有错误信息时，添加错误消息
			err = fmt.Errorf("查无此角色")
		}
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("role id : %v", in.RoleId),
			Error:     err,
			ErrorCode: "UR0193",
		})
		return nil, err
	}
	// Step 2: 更新内容
	role.Name = in.NewRoleName
	err = l.RoleModel.Update(l.ctx, role)
	if err != nil {
		log.Error(l.ctx, log.ErrorContent{
			Message:   in.NewRoleName,
			Error:     err,
			ErrorCode: "UR0194",
		})
		return nil, err
	}
	return &user_server.UpdateRoleResponse{
		RoleId: in.RoleId,
	}, nil
}
