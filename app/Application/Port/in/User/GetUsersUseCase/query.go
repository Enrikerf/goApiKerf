package GetUsersUseCase

import "github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Models"

type Query interface{
	Get() []Models.User
}