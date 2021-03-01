package Controllers

import (
	"fmt"
	"github.com/Enrikerf/goApiKerf/app/Config/Injection/Adapter"
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
	var getUserQuery = Adapter.GetUsersQuery()

	context.String(http.StatusOK, fmt.Sprintf(getUserQuery.Get()))
}


func postUsers(context *gin.Context) {
	context.String(http.StatusOK, fmt.Sprintf("creando user"))
}
