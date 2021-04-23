package Services

import (
	"github.com/Enrikerf/goApiKerf/app/Application/Port/out/Database/Users"
)

type GetUsersService struct {
	UsersPort Users.GetUsersPort
}

func (service GetUsersService) Get() string {
	var getUserPort = service.UsersPort
	return getUserPort.Get()
}
