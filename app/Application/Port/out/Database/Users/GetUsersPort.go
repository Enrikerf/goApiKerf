package Users

import (
	"github.com/Enrikerf/goApiKerf/app/Domain"
)

type GetUsersPort interface{
	GetUsers() []Domain.User
}