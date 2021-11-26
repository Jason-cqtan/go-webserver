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
		c.SysErrJson(&commonResponse{
			BizCode: 4,
			Msg: fmt.Sprintf("error:%v",err),
		})
		return
	}

	// 假设返回正确id
	c.OkJson(&commonResponse{
		BizCode: 1,
		Msg: "success",
		Data: 1,
	})

}



// 响应json格式
type commonResponse struct {
	BizCode int
	Msg string
	Data interface{}
}

func main() {
	server := web.NewServer("openresty")

	server.Route("/signup",SignUp)
	server.Start(":8099")
}
