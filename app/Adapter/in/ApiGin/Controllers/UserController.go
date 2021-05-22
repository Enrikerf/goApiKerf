package Controllers

import (
	"fmt"
	"github.com/Enrikerf/goApiKerf/app/Application/Port/in/User/CreateUser"
	"github.com/Enrikerf/goApiKerf/app/Application/Port/in/User/GetUsers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	GetUsersUseCase   GetUsers.Query
	CreateUserUseCase CreateUser.UseCase
}

func (userController UserController) LoadUserControllerEndpoints(router *gin.Engine) {
	v1 := router.Group("/users")
	{
		v1.GET("", userController.getUsers)
		v1.POST("", userController.postUsers)
	}
}

func (userController UserController) getUsers(context *gin.Context) {
	var users = userController.GetUsersUseCase.GetUsersQuery()
	fmt.Println("{}", users)
	context.JSON(http.StatusOK, users)
}

func (userController UserController) postUsers(context *gin.Context) {
	var command = CreateUser.Command{
		Nickname: "new",
		Email:    "new",
		Password: "new",
	}
	user, err := userController.CreateUserUseCase.CreateUserUseCase(command)
	if err != nil {
		context.JSON(http.StatusInternalServerError, "")
	}
	context.JSON(http.StatusOK, user)
}
