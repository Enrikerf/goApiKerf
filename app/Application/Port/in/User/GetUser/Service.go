package GetUser

import (
	"github.com/Enrikerf/goApiKerf/app/Application/Port/out/Database/Users"
	"github.com/Enrikerf/goApiKerf/app/Domain"
)

type Service struct {
	UsersPort Users.GetUserPort
}

func (service Service) GetUserQuery(id string) Domain.User {
	return service.UsersPort.GetUser(id)
}