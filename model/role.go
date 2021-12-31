// Package model
// @Author: Nick Bi
// @Date:
// @Desc: todo...
///**
// Package model
// @Author: Nick Bi
// @Date: 2021-12-31
// @Desc: User model
///**
package model

import "time"

type Role struct {
	ID         int    `json "id"`
	Name       string `json "name"`
	Type       string
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
func (Role) TableName() string {
	return "t_role"
}
