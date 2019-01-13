package router

import (
	"realgo"
	"myframe/controllers"
	"myframe/middleware"
)

func Init(server *realgo.WebServer)  {
	server.EUSE(middleware.PrintLog)
	server.GET("/index", controllers.Index)
}



