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

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoiaHVhbmdoYWkiLCJwYXNzd29yZCI6IjEyMzQ1NiIsImV4cCI6MTY2MDkwMDEzNywiaXNzIjoidG9kb19saXN0In0.RkV3BY3zIy1OvY6zMK6Sxs03ToN7YhtbuuySUXCi2EQ
