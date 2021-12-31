// Package router
// @Author: Nick Bi
// @Date:
// @Desc: todo...
///**
package router

import (
	"github.com/gin-gonic/gin"
)

func init() {

	r := gin.Default()
	RegisterUser(r)
	r.Run(":8080")
}
