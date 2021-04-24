package Routes

import (
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/ApiGin/Controllers"
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Adapter"
	"github.com/Enrikerf/goApiKerf/app/Application/Services"
	"github.com/gin-gonic/gin"
)

func ConfigUserRoutes(route *gin.Engine)  {
	var getUserAdapter = Adapter.GetUsersAdapter{}
	var service = Services.GetUsersService{UsersPort: getUserAdapter}
	var userController = Controllers.UserController{GetUsersUseCase: service}
	userController.LoadUserControllerEndpoints(route)
}