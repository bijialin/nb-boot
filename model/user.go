// Package model
// @Author: Nick Bi
// @Date: 2021-12-31
// @Desc: User model
///**
package model

import "time"

type User struct {
	ID         int    `json "id"`
	Username   string `json "username"`
	Password   string
	Phone      string
	Avatar     string
	State      int
	CreateTime time.Time
	CreateBy   int
	UpdateTime time.Time
	UpdateBy   int
}

// TableName
//  @Description: 表名与实体名不一致，实体中定义表名方法返回给gorm使用
//  @receiver User
//  @return string
//
func (User) TableName() string {
	return "t_user"
}
