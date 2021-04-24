package Services

import (
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Models"
	"github.com/Enrikerf/goApiKerf/app/Application/Port/out/Database/Users"
)

type GetUsersService struct {
	UsersPort Users.GetUsersPort
}

func (service GetUsersService) Get() []Models.User {
	var getUserPort = service.UsersPort
	return getUserPort.Get()
}
