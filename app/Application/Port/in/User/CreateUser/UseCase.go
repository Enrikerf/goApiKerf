package CreateUser

import "github.com/Enrikerf/goApiKerf/app/Domain"

type UseCase interface{
	CreateUserUseCase(command Command) (Domain.User, error)
}