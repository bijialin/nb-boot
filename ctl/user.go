// Package ctl
// @Author: Nick Bi
// @Date: 2021-12-31
// @Desc: todo...
///**
package ctl

import (
	Res "github.com/bijialin/nb-boot/common/res"
	"github.com/bijialin/nb-boot/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserCtl struct {
	svc service.UserSvc
}

func NewUserCtlImpl() *UserCtl {
	return &UserCtl{
		svc: service.NewUserSvcImpl(),
	}
}
func (a *UserCtl) Get(c *gin.Context) {
	c.JSON(200, Res.Ok())
}
func (a *UserCtl) AddUser(c *gin.Context) {

}
func (a *UserCtl) UpdateUser(c *gin.Context) {

}
func (a *UserCtl) List(c *gin.Context) {

}
func (a *UserCtl) DeleteUser(c *gin.Context) {

}

func (a *UserCtl) Test(c *gin.Context) {
	c.JSON(http.StatusOK, Res.Ok())
}
