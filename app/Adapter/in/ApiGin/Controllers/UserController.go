package Controllers

import (
	"fmt"
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Adapter"
	"github.com/Enrikerf/goApiKerf/app/Application/Services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadUserControllerEndpoints(router *gin.Engine) {
	v1 := router.Group("/users")
	{
		v1.GET("", getUsers)
		v1.POST("", postUsers)
	}
}


func getUsers(context *gin.Context) {
	var getUserAdapter = Adapter.GetUsersAdapter{}
	var service = Services.GetUsersService{UsersPort: getUserAdapter}

	context.String(http.StatusOK, fmt.Sprintf(service.Get()))
}


func postUsers(context *gin.Context) {
	context.String(http.StatusOK, fmt.Sprintf("creando user"))
}
