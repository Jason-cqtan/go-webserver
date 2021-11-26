package main

import (
	"github.com/Jason-cqtan/webserver/demo"
	web "github.com/Jason-cqtan/webserver/lib"
)


func main() {
	server := web.NewServer("openresty")
	server.Route("POST","/signup",demo.SignUp)
	server.Start(":8099")
}
