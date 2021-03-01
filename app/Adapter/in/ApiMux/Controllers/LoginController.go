package Controllers

import (
	"encoding/json"
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/ApiMux/Auth"
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/ApiMux/Responses"
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/ApiMux/Utils/FormatError"
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Models"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := Models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := FormatError.FormatError(err.Error())
		Responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	Responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := Models.User{}

	err = server.DB.Debug().Model(Models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = Models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return Auth.CreateToken(user.ID)
}
