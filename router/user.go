// Package router
// @Author: Nick Bi
// @Date:
// @Desc: todo...
///**
package router

import (
	"github.com/bijialin/nb-boot/common/global"
	"github.com/bijialin/nb-boot/ctl"
	"github.com/gin-gonic/gin"
)

func RegisterUser(router *gin.Engine) {
	global.Log.Info("start register user router")
	useCtl := ctl.NewUserCtlImpl()
	v := router.Group("api/v1/user")
	{
		v.GET("/:id", useCtl.Get)
		v.GET("/test", useCtl.Test)

	}
}
