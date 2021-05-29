package Adapter

import (
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Models"
	"github.com/Enrikerf/goApiKerf/app/Domain"
	"gorm.io/gorm"
)

type UsersAdapter struct {
	Orm *gorm.DB
}

func (userAdapter UsersAdapter) GetUsers() []Domain.User {
	var users []Models.User
	userAdapter.Orm.Find(&users)
	var domainUsers []Domain.User
	return domainUsers
}

func (userAdapter UsersAdapter) SaveUser(user Domain.User) (Domain.User,error) {
	var userEntity  Models.User
	userEntity.Nickname = user.Nickname
	userEntity.Email = user.Nickname
	userEntity.Password = user.Password
	saveUser, err := userEntity.SaveUser(userAdapter.Orm)
	if err != nil {
		return user, err
	}
	user.Id = saveUser.ID
	//userAdapter.Orm.Create(&userEntity)
	return user, nil
}
