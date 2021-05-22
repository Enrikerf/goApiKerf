package Routes

import (
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/ApiGin/Controllers"
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Adapter"
	"github.com/Enrikerf/goApiKerf/app/Application/Port/in/User/GetUsers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigUserRoutes(route *gin.Engine, db *gorm.DB) {
	var userAdapter = Adapter.UsersAdapter{Orm: db}
	var service = GetUsers.Service{UsersPort: userAdapter}
	var userController = Controllers.UserController{GetUsersUseCase: service}
	userController.LoadUserControllerEndpoints(route)
}
