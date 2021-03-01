package ApiGin

import (
	"fmt"
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/ApiGin/Config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServerGin struct {
	Engine * gin.Engine
 }

func (serverGin *ServerGin) Run(){
		serverGin.Engine = gin.Default()
		Config.Routes(serverGin.Engine)
		serverGin.Engine.Run(":8081")
}

func pong(context *gin.Context) {
	context.String(http.StatusOK, fmt.Sprintf("files uploaded!"))
}


