package GetUsers

import (
	"github.com/Enrikerf/goApiKerf/app/Domain"
)

type Query interface{
	GetUsersQuery() []Domain.User
}