package Routes

import (
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/ApiGin/Controllers"
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Adapter"
	"github.com/Enrikerf/goApiKerf/app/Application/Port/in/User/CreateUser"
	"github.com/Enrikerf/goApiKerf/app/Application/Port/in/User/GetUser"
	"github.com/Enrikerf/goApiKerf/app/Application/Port/in/User/GetUsers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigUserRoutes(route *gin.Engine, db *gorm.DB) {
	var userAdapter = Adapter.UsersAdapter{Orm: db}
	var GetUserService = GetUser.Service{UsersPort: userAdapter}
	var GetUsersService = GetUsers.Service{UsersPort: userAdapter}
	var CreateUserService = CreateUser.Service{SaveUserPort: userAdapter}
	var userController = Controllers.UserController{
		GetUserUseCase:    GetUserService,
		GetUsersUseCase:   GetUsersService,
		CreateUserUseCase: CreateUserService,
	}
	userController.LoadUserControllerEndpoints(route)
}
