package Adapter

import (
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Models"
	"gorm.io/gorm"
)

type GetUsersAdapter struct {
	Orm *gorm.DB
}

func (service GetUsersAdapter) Get() []Models.User {
	var users []Models.User
	service.Orm.Find(&users)
	return users
}
