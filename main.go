package main

import (
	"memo_gin/conf"
	"memo_gin/routes"
)

func main() {
	conf.Init()
	e := routes.NewRouter()
	e.Run(conf.HttpPort) //HttpPort=3000
}
