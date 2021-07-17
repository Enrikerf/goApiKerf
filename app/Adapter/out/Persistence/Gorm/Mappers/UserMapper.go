package Mappers

import (
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Models"
	"github.com/Enrikerf/goApiKerf/app/Domain"
)

func Persistence2Domain(modelUser Models.User) (Domain.User, error) {
	var user = Domain.User{}
	user.Id = modelUser.ID
	user.Email = modelUser.Email
	return user, nil
}
func Persistence2DomainArray(modelUsers []Models.User) ([]Domain.User, error) {
	var domainUsers []Domain.User
	for _, modelUser := range modelUsers {
		var user = Domain.User{}
		user.Id = modelUser.ID
		user.Email = modelUser.Email
		domainUsers = append(domainUsers, user)
	}
	return domainUsers, nil
}
