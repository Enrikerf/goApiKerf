package CreateUser

import (
	"github.com/Enrikerf/goApiKerf/app/Application/Port/out/Database/Users"
	"github.com/Enrikerf/goApiKerf/app/Domain"
)

type Service struct {
	SaveUserPort Users.SaveUserPort
}

func (service Service) CreateUserUseCase(command Command) (Domain.User, error) {
	var user = Domain.User{}
	user.Password = command.Password
	user.Email = command.Email
	user.Nickname = command.Nickname
	return service.SaveUserPort.SaveUser(user)
}
