package userserverlogic

import (
	"context"
	"fmt"
	"github.com/zzp-Z/UserServer/db/crud"
	"github.com/zzp-Z/UserServer/db/enum"
	"github.com/zzp-Z/UserServer/internal/logic"
	"github.com/zzp-Z/UserServer/internal/svc"
	"github.com/zzp-Z/UserServer/logs"
	"github.com/zzp-Z/UserServer/user_server"
	"regexp"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	UserModel crud.UserModel
	Tools     *logic.Tools
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		UserModel: crud.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf),
		Tools:     &logic.Tools{},
	}
}

// Register 注册新用户 将用户信息写入数据库，应当在确认用户邮箱后调用
func (l *RegisterLogic) Register(in *user_server.RegisterRequest) (*user_server.RegisterResponse, error) {
	// STEP: 校验各种信息
	in, err := validateInfo(in)
	if err != nil {
		logs.Error(nil, logs.ErrorContent{
			Message:   in.Email,
			Error:     err,
			ErrorCode: "RR900",
		})
		return nil, err
	}
	// STEP: hash密码
	hashPassword, err := l.Tools.HashPassword(in.Password)
	if err != nil {
		logs.Error(nil, logs.ErrorContent{
			Message:   in.Email,
			Error:     err,
			ErrorCode: "RR901",
		})
		return nil, err
	}
	// STEP: 插入数据库
	insert, err := l.UserModel.Insert(l.ctx, &crud.User{
		Username:     in.Username,
		HashPassword: hashPassword,
		Email:        in.Email,
		Status:       enum.UserStatusEnabled.String(),
	})
	if err != nil {
		logs.Error(nil, logs.ErrorContent{
			Message:   in.Email,
			Error:     err,
			ErrorCode: "RR902",
		})
		return nil, err
	}
	// STEP: 返回用户ID
	id, err := insert.LastInsertId()
	if err != nil {
		logs.Error(nil, logs.ErrorContent{
			Message:   in.Email,
			Error:     err,
			ErrorCode: "RR903",
		})
		return nil, err
	}
	return &user_server.RegisterResponse{
		UserId: uint32(id),
	}, nil
}

// 校验各种信息
func validateInfo(in *user_server.RegisterRequest) (*user_server.RegisterRequest, error) {
	// 记录原始输入
	rawUsername := in.Username
	rawPassword := in.Password
	rawEmail := in.Email

	// 去除前后空格
	username := strings.TrimSpace(rawUsername)
	password := strings.TrimSpace(rawPassword)
	email := strings.TrimSpace(rawEmail)

	in.Username = username
	in.Password = password
	in.Email = email

	// 校验用户名
	if len(username) == 0 || len(username) > 20 {
		err := fmt.Errorf("用户名不能为空且长度不能超过20个字符")
		logs.Error(nil, logs.ErrorContent{
			Message:   rawUsername,
			Error:     err,
			ErrorCode: "RVI900",
		})
		return in, err
	}
	// 校验用户名中是否包含非法字符
	if containsInvalidChars(username) {
		err := fmt.Errorf("用户名包含非法字符")
		logs.Error(nil, logs.ErrorContent{
			Message:   rawUsername,
			Error:     err,
			ErrorCode: "RVI901",
		})
		return in, err
	}

	// 校验密码
	if len(rawPassword) > 0 && rawPassword != password {
		err := fmt.Errorf("密码前后含有空格")
		logs.Error(nil, logs.ErrorContent{
			Message:   rawPassword,
			Error:     err,
			ErrorCode: "RVI902",
		})
		return in, err
	}
	if len(password) == 0 {
		err := fmt.Errorf("密码不能为空")
		logs.Error(nil, logs.ErrorContent{
			Message:   rawPassword,
			Error:     err,
			ErrorCode: "RVI903",
		})
		return in, err
	}

	// 校验电子邮件
	if len(rawEmail) > 0 && rawEmail != email {
		err := fmt.Errorf("电子邮件前后含有空格")
		logs.Error(nil, logs.ErrorContent{
			Message:   rawEmail,
			Error:     err,
			ErrorCode: "RVI904",
		})
		return in, err
	}
	if len(email) == 0 {
		err := fmt.Errorf("电子邮件不能为空")
		logs.Error(nil, logs.ErrorContent{
			Message:   rawEmail,
			Error:     err,
			ErrorCode: "RVI905",
		})
		return in, err
	}
	if !isValidEmail(email) {
		err := fmt.Errorf("电子邮件格式不正确")
		logs.Error(nil, logs.ErrorContent{
			Message:   rawEmail,
			Error:     err,
			ErrorCode: "RVI906",
		})
		return in, err
	}
	return in, nil
}

// containsInvalidChars 检查字符串是否包含非法字符
func containsInvalidChars(s string) bool {
	invalidChars := []rune{'(', ')', ';', '=', '<', '>', ','}
	for _, char := range s {
		for _, invalidChar := range invalidChars {
			if char == invalidChar {
				return true
			}
		}
	}
	return false
}

// isValidEmail 校验电子邮件格式
func isValidEmail(email string) bool {
	// 正则表达式用于检查电子邮件格式
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
