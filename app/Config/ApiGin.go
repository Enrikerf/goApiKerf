package Config

import (
	"fmt"
	"github.com/Enrikerf/goApiKerf/app/Adapter/out/Persistence/Gorm/Models"
	"github.com/Enrikerf/goApiKerf/app/Config/Routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type ApiGin struct {
	DB         *gorm.DB
	Engine     *gin.Engine
	portNumber string
}

func (apiGin *ApiGin) Initialize(serverPortNumber, dbUser, dbPassword, dbPort, dbHost, dbName string) {
	apiGin.portNumber = serverPortNumber
	apiGin.Engine = gin.Default()
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dbUrl, // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Models.User{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to migrate database")
	}
	Routes.ConfigUserRoutes(apiGin.Engine, db)
}

func (apiGin *ApiGin) Run() {
	fmt.Printf("listening on portNumber: %s", apiGin.portNumber)
	log.Fatal(apiGin.Engine.Run(apiGin.portNumber))
}
