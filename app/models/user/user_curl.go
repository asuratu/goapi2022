// Package user 存放用户 Model 相关逻辑
// 方法的接收者为 (userModel *User)
package user

import (
	"gohub/pkg/database"
	"gohub/pkg/hash"
)

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}
