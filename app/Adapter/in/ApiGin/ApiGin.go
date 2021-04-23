package ApiGin

import (
	"fmt"
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/ApiGin/Config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiGin struct {
	Engine * gin.Engine
 }

func (apiGin *ApiGin) Run(){
		apiGin.Engine = gin.Default()
		Config.Routes(apiGin.Engine)
		apiGin.Engine.Run(":8081")
}

func pong(context *gin.Context) {
	context.String(http.StatusOK, fmt.Sprintf("files uploaded!"))
}


