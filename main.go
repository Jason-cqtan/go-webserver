package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Server interface {
	Route(pattern string,handlerFunc http.HandlerFunc)
	Start(addr string) error
}

type sdkHttpServer struct {
	name string
}

func (s *sdkHttpServer) Route(pattern string,handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern,handlerFunc)
}

func (s *sdkHttpServer) Start(addr string) error {
	fmt.Printf("listening %s\n",addr)
	return http.ListenAndServe(addr,nil)
}


type signUpReq struct {
	Email string `json:"email"`
	Password string `json:"password"`
	confirmedPassword string `json:"confirmed_password"`
}

func SignUp(w http.ResponseWriter,r *http.Request) {
	req := &signUpReq{}
	c := &Context{w,r}
	err := c.ReadJson(req)
	if err != nil {
		fmt.Fprintf(w,"error:%v",err)
		return
	}

	// 假设返回正确id
	fmt.Fprintf(w,"id:%d",1)

}

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

// ReadJson 读取body数据
func (c *Context) ReadJson(data interface{}) error {
	body,err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body,data)
}


func main() {
	server := &sdkHttpServer{
		name: "openresty",
	}

	server.Route("/signup",SignUp)
	server.Start(":8099")
}
