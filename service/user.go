package service

import (
	"github.com/bijialin/nb-boot/common/global"
	"github.com/bijialin/nb-boot/model"
	"gorm.io/gorm"
)

type UserSvc interface {
	GetUser(req model.User)
	GetUserList(req model.User)
	Add(req model.User)
	Update(req model.User)
}

//
// UserSvcImpl
//  @Description:  service的实现定义，可以使用全局注册的db进行操作，减少dao层定义，具体按项目分层架构
//
type UserSvcImpl struct {
	db *gorm.DB
}

func NewUserSvcImpl() UserSvc {
	return &UserSvcImpl{
		db: global.Db,
	}
}

func (s *UserSvcImpl) GetUser(req model.User) {
	//TODO implement me
	panic("implement me")
}

func (s *UserSvcImpl) GetUserList(req model.User) {
	//TODO implement me
	panic("implement me")
}

func (s *UserSvcImpl) Add(req model.User) {
	s.db.Save(req)
}

func (s UserSvcImpl) Update(req model.User) {
	s.db.Updates(req)
}

func (s UserSvcImpl) Delete(id int) {
	s.db.Delete(id)
}
