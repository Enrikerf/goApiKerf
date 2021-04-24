package Users

import "github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Models"

type GetUsersPort interface{
	Get() []Models.User
}