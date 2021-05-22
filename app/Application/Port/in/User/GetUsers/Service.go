package GetUsers

import (
	"github.com/Enrikerf/goApiKerf/app/Application/Port/out/Database/Users"
	"github.com/Enrikerf/goApiKerf/app/Domain"
)

type Service struct {
	UsersPort Users.GetUsersPort
}

func (service Service) GetUsersQuery() []Domain.User {
	return service.UsersPort.GetUsers()
}
