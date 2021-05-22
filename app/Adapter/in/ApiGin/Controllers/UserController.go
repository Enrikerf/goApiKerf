package Controllers

import (
	"fmt"
	"github.com/Enrikerf/goApiKerf/app/Application/Port/in/User/GetUsers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	GetUsersUseCase GetUsers.Query
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
	context.String(http.StatusOK, fmt.Sprintf("creando user"))
}
