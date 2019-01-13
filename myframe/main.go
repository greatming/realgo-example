package main

import (
	"myframe/router"
	"realgo"
	"myframe/config"
)

func init()  {
	config.Init()
}


func main()  {
	app := realgo.New()
	app.RegisterRouter(router.Init)
	app.StartServer()

}
