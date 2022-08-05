// Package user 存放用户 Model 相关逻辑
// 本文件处理 user 的数据库相关操作
// 方法的接收者为 (userModel *User)
package user

import (
	"gohub/pkg/database"
)

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}
