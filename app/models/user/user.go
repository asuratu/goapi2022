// Package user 存放用户 Model 相关逻辑
package user

import (
    "gohub/app/models"
)

// User 用户模型
type User struct {
    models.BaseModel

    Name string `json:"name,omitempty"`

    City         string `json:"city,omitempty"`
    Introduction string `json:"introduction,omitempty"`
    Avatar       string `json:"avatar,omitempty"`

    Email    string `json:"-"`
    Phone    string `json:"-"`
    Password string `json:"-"`

    models.CommonTimestampsField
}
