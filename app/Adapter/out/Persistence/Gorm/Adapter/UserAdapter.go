package Adapter

import (
	"fmt"
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Mappers"
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Models"
	"github.com/Enrikerf/goApiKerf/app/Domain"
	"gorm.io/gorm"
)

type UsersAdapter struct {
	Orm *gorm.DB
}

func (userAdapter UsersAdapter) GetUser(id string) Domain.User {
	var user Models.User
	var userEntity Models.User
	user,_ = userEntity.FindUserByID(userAdapter.Orm,id)
	var domainUsers Domain.User
	domainUsers,_  = Mappers.Persistence2Domain(user)
	return domainUsers
}

func (userAdapter UsersAdapter) GetUsers() []Domain.User {
	var users []Models.User
	var userEntity Models.User
	users,_ = userEntity.FindAllUsers(userAdapter.Orm)
	fmt.Println("{}", users)
	var domainUsers []Domain.User
	domainUsers,_  = Mappers.Persistence2DomainArray(users)
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
	return user, nil
}
