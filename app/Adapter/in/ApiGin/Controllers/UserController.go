package Controllers

import (
	"github.com/Enrikerf/goApiKerf/app/Application/Port/in/User/CreateUser"
	"github.com/Enrikerf/goApiKerf/app/Application/Port/in/User/GetUser"
	"github.com/Enrikerf/goApiKerf/app/Application/Port/in/User/GetUsers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	GetUserUseCase    GetUser.Query
	GetUsersUseCase   GetUsers.Query
	CreateUserUseCase CreateUser.UseCase
}

func (userController UserController) LoadUserControllerEndpoints(router *gin.Engine) {
	v1 := router.Group("/users")
	{
		v1.GET("/:id", userController.getUser)
		v1.GET("", userController.getUsers)
		v1.POST("", userController.postUsers)
	}
}

func (userController UserController) getUser(context *gin.Context) {
	id := context.Param("id")
	var user = userController.GetUserUseCase.GetUserQuery(id)
	context.JSON(http.StatusOK, user)
}

func (userController UserController) getUsers(context *gin.Context) {
	var users = userController.GetUsersUseCase.GetUsersQuery()
	context.JSON(http.StatusOK, users)
}

func (userController UserController) postUsers(context *gin.Context) {
	var command CreateUser.Command
	err2 := context.BindJSON(&command)
	if err2 != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, err := userController.CreateUserUseCase.CreateUserUseCase(command)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	}
	context.JSON(http.StatusOK, user)
}
