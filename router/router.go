// Package router
// @Author: Nick Bi
// @Date:
// @Desc: todo...
///**
package router

import (
	"github.com/bijialin/nb-boot/common/global"
	_ "github.com/bijialin/nb-boot/docs" //必须导入生成的文档包，不然会读取不到文档
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	//默认服务端口
	defaultServerPort = "8080"
)

// @title nb-boot api
// @version 1.0.0
// @contact.name Nick Bi
// @contact.url http://www.swagger.io/support
// @contact.email
// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/
func init() {
	global.Log.Info("start init routers")
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	RegisterUser(r)
	port := global.AppConfig.Server.Port
	if port != "" {
		port = defaultServerPort
	}
	r.Run(":" + port)
}
