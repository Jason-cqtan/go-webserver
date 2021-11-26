package lib

import (
	"fmt"
	"net/http"
)

type Server interface {
	Route(pattern string,handlerFunc func(c *Context))
	Start(addr string) error
}

type sdkHttpServer struct {
	Name string
}

// 响应json格式
type commonResponse struct {
	BizCode int
	Msg string
	Data interface{}
}

func (s *sdkHttpServer) Route(pattern string,handlerFunc func(c *Context)) {
	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		handlerFunc(&Context{
			W: writer,
			R: request,
		})
	})
}

func (s *sdkHttpServer) Start(addr string) error {
	fmt.Printf("listening %s\n",addr)
	return http.ListenAndServe(addr,nil)
}

// NewServer 新建实例
func NewServer(name string) *sdkHttpServer{
	return &sdkHttpServer{
		Name: name,
	}
}

func NewRes(code int,msg string,data interface{}) *commonResponse{
	return &commonResponse{
		BizCode: code,
		Msg: msg,
		Data: data,
	}
}

