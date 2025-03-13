package rolelogic

import (
	"context"
	"fmt"
	"github.com/zzp-Z/UserServer/db/crud"
	"github.com/zzp-Z/UserServer/log"
	"time"

	"github.com/zzp-Z/UserServer/UserService"
	"github.com/zzp-Z/UserServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RoleModel crud.RoleModel
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		RoleModel: crud.NewRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// DeleteRole 删除角色
func (l *DeleteRoleLogic) DeleteRole(in *UserService.DeleteRoleRequest) (*UserService.DeleteRoleResponse, error) {
	// Step 1: 查询角色
	role, err := l.RoleModel.FindOne(l.ctx, in.RoleId)
	if err != nil || role.DeletedAt.IsZero() != true {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("role id : %v", in.RoleId),
			Error:     err,
			ErrorCode: "DR0193",
		})
		err = fmt.Errorf("查无此角色")
		return nil, err
	}
	// Step 2: 删除角色
	role.DeletedAt = time.Now()
	err = l.RoleModel.Update(l.ctx, role)
	if err != nil {
		log.Error(nil, log.ErrorContent{
			Message:   fmt.Sprintf("role id : %v", in.RoleId),
			Error:     err,
			ErrorCode: "DR0194",
		})
		return nil, err
	}

	return &UserService.DeleteRoleResponse{
		RoleId: role.Id,
	}, nil
}
