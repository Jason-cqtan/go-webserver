package main

import (
	"fmt"
	web "github.com/Jason-cqtan/webserver/lib"

)

type signUpReq struct {
	Email string `json:"email"`
	Password string `json:"password"`
	confirmedPassword string `json:"confirmed_password"`
}

func SignUp(c *web.Context) {
	req := &signUpReq{}
	err := c.ReadJson(req)
	if err != nil {
		c.SysErrJson(web.NewRes(500,fmt.Sprintf("error:%v",err),nil))
		return
	}

	// 假设返回正确id
	c.OkJson(web.NewRes(200,"success",1))
}

func main() {
	server := web.NewServer("openresty")
	server.Route("/signup",SignUp)
	server.Start(":8099")
}
