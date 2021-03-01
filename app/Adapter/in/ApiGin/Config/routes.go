package Config

import (
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/ApiGin/Controllers"
	"github.com/gin-gonic/gin"
)

func Routes(route * gin.Engine) {
	Controllers.LoadUserControllerEndpoints(route)
}
