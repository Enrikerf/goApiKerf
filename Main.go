package main

import (
	"github.com/Enrikerf/goApiKerf/app/Config"
)

func main() {
	//ApiMux.Run()
	var ginServer = Config.GinServer{}
	ginServer.Run()
}
