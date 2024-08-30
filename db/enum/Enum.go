package enum

type UserStatus string

const (
	UserStatusDisabled UserStatus = "0" // 禁用
	UserStatusEnabled  UserStatus = "1" // 启用
)

func (u UserStatus) String() string {
	return string(u)
}
