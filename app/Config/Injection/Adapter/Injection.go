package Adapter

import (
	"github.com/Enrikerf/goApiKerf/app/Application/Port/in/User/GetUsers"
	"github.com/Enrikerf/goApiKerf/app/Application/Services"
	"github.com/google/wire"
)

func GetUsersQuery() Services.GetUsersService {
	wire.Build(wire.Bind(new(Services.GetUsersService), new(GetUsers.Query)))
	return Services.GetUsersService{}
}
