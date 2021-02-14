package Api

import (
	"fmt"
	"github.com/Enrikerf/goApiKerf/app/Adapter/in/Api/Controllers"
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Seeds"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var server = Controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	Seeds.Load(server.DB)

	server.Run(":8080")

}
