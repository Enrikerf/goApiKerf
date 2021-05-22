package Adapter

import (
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Models"
	"github.com/Enrikerf/goApiKerf/app/Domain"
	"gorm.io/gorm"
)

type UsersAdapter struct {
	Orm *gorm.DB
}

func (service UsersAdapter) GetUsers() []Domain.User {
	var users []Models.User
	service.Orm.Find(&users)
	var domainUsers []Domain.User
	return domainUsers
}

func (service UsersAdapter) SaveUser(user Domain.User) (Domain.User,error) {
	return user,nil
}
