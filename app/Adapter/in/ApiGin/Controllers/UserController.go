package Controllers

import (
	"fmt"
	"github.com/Enrikerf/goApiKerf/app/Application/Port/in/User/GetUsersUseCase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	GetUsersUseCase GetUsersUseCase.Query
}

func (userController UserController) LoadUserControllerEndpoints(router *gin.Engine) {
	v1 := router.Group("/users")
	{
		v1.GET("", userController.getUsers)
		v1.POST("", userController.postUsers)
	}
}

func (userController UserController) getUsers(context *gin.Context) {
	context.String(http.StatusOK, fmt.Sprintf(userController.GetUsersUseCase.Get()))
}

func (userController UserController) postUsers(context *gin.Context) {
	context.String(http.StatusOK, fmt.Sprintf("creando user"))
}
