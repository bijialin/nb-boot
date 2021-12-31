// Package router
// @Author: Nick Bi
// @Date:
// @Desc: todo...
///**
package router

import (
	"github.com/bijialin/nb-boot/common/global"
	"github.com/gin-gonic/gin"
)

func init() {
	global.Log.Info("start init routers")
	r := gin.Default()
	RegisterUser(r)
	r.Run(":8080")
}
