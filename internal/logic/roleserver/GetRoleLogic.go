package roleserverlogic

import (
	"context"
	"github.com/zzp-Z/UserServer/db/crud"
	"github.com/zzp-Z/UserServer/log"
	"strconv"

	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	RoleModel crud.RoleModel
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		RoleModel: crud.NewRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// GetRole 获取角色信息
func (l *GetRoleLogic) GetRole(in *user_server.GetRoleRequest) (*user_server.GetRoleResponse, error) {
	// Step 1: 查询角色
	role, err := l.RoleModel.FindOne(l.ctx, in.RoleId)
	if err != nil {
		log.Error(l.ctx, log.ErrorContent{
			Message:   strconv.FormatUint(in.RoleId, 10),
			Error:     err,
			ErrorCode: "GR01311",
		})
		return nil, err
	}

	return &user_server.GetRoleResponse{
		RoleId:   role.Id,
		RoleName: role.Name,
	}, nil
}
