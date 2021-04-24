package main

import (
	"github.com/Enrikerf/goApiKerf/app/Config"
)

func main() {
	//ApiMux.Run()
	var apiGin = Config.ApiGin{}
	apiGin.Run()
}
