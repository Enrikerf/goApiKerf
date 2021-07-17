package Users

import (
	"github.com/Enrikerf/goApiKerf/app/Domain"
)

type SaveUserPort interface{
	SaveUser(user Domain.User) (Domain.User,error)
}