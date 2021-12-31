// Package global
// @Author: Nick Bi
// @Date:
// @Desc: todo...
///**
package global

import (
	"github.com/bijialin/nb-boot/common/base"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Log       *zap.SugaredLogger //日志实例
	Db        *gorm.DB           //数据库实例
	AppConfig *base.Config       //项目配置实例
)
