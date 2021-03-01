package Middlewares

import (
	"errors"
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/ApiMux/Auth"
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/ApiMux/Responses"
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := Auth.TokenValid(r)
		if err != nil {
			Responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
