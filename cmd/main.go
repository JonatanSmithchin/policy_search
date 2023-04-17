package main

import (
	"police_search/Router"
)

func main() {

	//gin.DisableConsoleColor()
	//
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//configs.InitConfig(".")
	Router.InitRouter()

}
