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

// WriteJson 返回json和响应封装
func (c *Context) WriteJson(code int,data interface{}) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.W.Write(bs)
	if err != nil {
		return err
	}
	c.W.Header().Set("Content-type:application/json","charset=utf-8")
	c.W.WriteHeader(code)
	return nil
}

// OkJson 辅助返回
func (c *Context) OkJson(data interface{}) error {
	return c.WriteJson(http.StatusOK,data)
}

func (c *Context) SysErrJson(data interface{}) error {
	return c.WriteJson(http.StatusInternalServerError,data)
}

func (c *Context) BadRequestJson(data interface{}) error {
	return c.WriteJson(http.StatusBadRequest,data)
}

// 响应json格式
type commonResponse struct {
	BizCode int
	Msg string
	Data interface{}
}

func main() {
	server := &sdkHttpServer{
		name: "openresty",
	}

	server.Route("/signup",SignUp)
	server.Start(":8099")
}
