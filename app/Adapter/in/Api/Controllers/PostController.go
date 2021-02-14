package Controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/Api/Auth"
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/Api/Responses"
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/Api/Utils/FormatError"
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := Models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post.Prepare()
	err = post.Validate()
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := Auth.ExtractTokenID(r)
	if err != nil {
		Responses.ERROR(w, http.StatusUnauthorized, errors.New("UnAuthorized"))
		return
	}
	if uid != post.AuthorID {
		Responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	postCreated, err := post.SavePost(server.DB)
	if err != nil {
		formattedError := FormatError.FormatError(err.Error())
		Responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, postCreated.ID))
	Responses.JSON(w, http.StatusCreated, postCreated)
}

func (server *Server) GetPosts(w http.ResponseWriter, r *http.Request) {

	post := Models.Post{}

	posts, err := post.FindAllPosts(server.DB)
	if err != nil {
		Responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	Responses.JSON(w, http.StatusOK, posts)
}

func (server *Server) GetPost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		Responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	post := Models.Post{}

	postReceived, err := post.FindPostByID(server.DB, pid)
	if err != nil {
		Responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	Responses.JSON(w, http.StatusOK, postReceived)
}

func (server *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the post id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		Responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the Auth token is valid and  get the user id from it
	uid, err := Auth.ExtractTokenID(r)
	if err != nil {
		Responses.ERROR(w, http.StatusUnauthorized, errors.New("UnAuthorized"))
		return
	}

	// Check if the post exist
	post := Models.Post{}
	err = server.DB.Debug().Model(Models.Post{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		Responses.ERROR(w, http.StatusNotFound, errors.New("Post not found"))
		return
	}

	// If a user attempt to update a post not belonging to him
	if uid != post.AuthorID {
		Responses.ERROR(w, http.StatusUnauthorized, errors.New("UnAuthorized"))
		return
	}
	// Read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	postUpdate := Models.Post{}
	err = json.Unmarshal(body, &postUpdate)
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != postUpdate.AuthorID {
		Responses.ERROR(w, http.StatusUnauthorized, errors.New("UnAuthorized"))
		return
	}

	postUpdate.Prepare()
	err = postUpdate.Validate()
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	postUpdate.ID = post.ID //this is important to tell the model the post id to update, the other update field are set above

	postUpdated, err := postUpdate.UpdateAPost(server.DB)

	if err != nil {
		formattedError := FormatError.FormatError(err.Error())
		Responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	Responses.JSON(w, http.StatusOK, postUpdated)
}

func (server *Server) DeletePost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid post id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		Responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user Authenticated?
	uid, err := Auth.ExtractTokenID(r)
	if err != nil {
		Responses.ERROR(w, http.StatusUnauthorized, errors.New("UnAuthorized"))
		return
	}

	// Check if the post exist
	post := Models.Post{}
	err = server.DB.Debug().Model(Models.Post{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		Responses.ERROR(w, http.StatusNotFound, errors.New("UnAuthorized"))
		return
	}

	// Is the Authenticated user, the owner of this post?
	if uid != post.AuthorID {
		Responses.ERROR(w, http.StatusUnauthorized, errors.New("UnAuthorized"))
		return
	}
	_, err = post.DeleteAPost(server.DB, pid, uid)
	if err != nil {
		Responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	Responses.JSON(w, http.StatusNoContent, "")
}
