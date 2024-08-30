package logic

import "golang.org/x/crypto/bcrypt"

type Tools struct{}

// HashPassword 生成哈希密码
func (t *Tools) HashPassword(password string) (hashedPassword string, err error) {
	byteHashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPassword = string(byteHashedPassword)
	return
}

// CheckPassword 验证密码
func (t *Tools) CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
