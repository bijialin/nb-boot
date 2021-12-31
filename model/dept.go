// Package model
// @Author: Nick Bi
// @Date:
// @Desc: todo...
///**
package model

import "time"

type Dept struct {
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
func (Dept) TableName() string {
	return "t_dept"
}
