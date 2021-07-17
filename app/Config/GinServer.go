package Config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type GinServer struct {
	ApiGin ApiGin
}

func (server *GinServer) Run() {

	var err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.ApiGin = ApiGin{}
	server.ApiGin.Initialize(
		os.Getenv("SERVER_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	server.ApiGin.Run()
}
