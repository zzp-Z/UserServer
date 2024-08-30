package model

import "gorm.io/gorm"

type UserStatus string

const (
	UserStatusDisabled UserStatus = "0" // 禁用
	UserStatusEnabled  UserStatus = "1" // 启用
)

type User struct {
	gorm.Model
	Username     string     `gorm:"unique;size:20;not null"`                       // 用户名，唯一，长度限制为20
	HashPassword string     `gorm:"not null;size:200"`                             // 密码 不能是空的
	Bio          string     `gorm:"size:200"`                                      // 简介 长度限制为200
	Message      string     `gorm:"size:200"`                                      // 消息 长度限制为200
	Email        string     `gorm:"unique;not null;size:200"`                      // 邮箱，唯一 不能是空的, 长度限制为200
	MoodId       int        `gorm:"not null;default:1"`                            // 心情，默认是1
	Status       UserStatus `gorm:"type:enum('0', '1');default:'0'" json:"status"` // 状态
}

// TableName 返回 User 结构体对应的表名。
func (User) TableName() string {
	return "user" // 设置表名为 "user"
}
