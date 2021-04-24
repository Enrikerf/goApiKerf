package Config

import (
	"github.com/Enrikerf/goApiKerf/app/Config/Routes"
	"github.com/gin-gonic/gin"
)

type ApiGin struct {
	Engine * gin.Engine
}

func (apiGin *ApiGin) Run(){
	apiGin.Engine = gin.Default()
	Routes.ConfigUserRoutes(apiGin.Engine)
	apiGin.Engine.Run(":8081")
}

