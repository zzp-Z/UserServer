package roleserverlogic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zzp-Z/UserServer/db/crud"
	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/log"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RoleModel crud.RoleModel
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		RoleModel: crud.NewRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// CreateRole 创建角色
func (l *CreateRoleLogic) CreateRole(in *user_server.CreateRoleRequest) (*user_server.CreateRoleResponse, error) {
	// Step 1: 查询角色是否存在
	_, err := l.RoleModel.FindOneByName(l.ctx, in.RoleName)
	if err == nil {
		err = fmt.Errorf("角色已存在")
		log.Error(l.ctx, log.ErrorContent{
			Message:   in.RoleName,
			Error:     err,
			ErrorCode: "CR0133",
		})
		return nil, err
	}

	// Step 2: 创建角色
	insert, err := l.RoleModel.Insert(l.ctx, &crud.Role{
		Name: in.RoleName,
		Description: sql.NullString{
			String: in.Description,
			Valid:  true,
		},
	})
	if err != nil {
		log.Error(l.ctx, log.ErrorContent{
			Message:   fmt.Sprintf("role info %v", in),
			Error:     err,
			ErrorCode: "CR0193",
		})
		return nil, err
	}
	id, _ := insert.LastInsertId()
	return &user_server.CreateRoleResponse{
		RoleId:   uint64(id),
		RoleName: in.RoleName,
	}, nil
}
