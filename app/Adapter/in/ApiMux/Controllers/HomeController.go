package Controllers

import (
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/ApiMux/Responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	Responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
