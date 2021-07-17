package Users

import (
	"github.com/Enrikerf/goApiKerf/app/Domain"
)

type GetUserPort interface{
	GetUser(id string) Domain.User
}