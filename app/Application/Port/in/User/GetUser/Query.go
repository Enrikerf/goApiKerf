package GetUser

import "github.com/Enrikerf/goApiKerf/app/Domain"

type Query interface{
	GetUserQuery(id string) Domain.User
}