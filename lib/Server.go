package lib

import (
	"fmt"
	"net/http"
)

type Routable interface {
	Route(method string,pattern string,handlerFunc func(c *Context))
}

// 组合 Routable
type Server interface {
	Routable
	Start(addr string) error
}

type Handler interface {
	http.Handler
	Routable
}

type sdkHttpServer struct {
	Name string
	handler Handler
}

// 响应json格式
type commonResponse struct {
	BizCode int
	Msg string
	Data interface{}
}


// ["":""]
type HandleBaseOnMap struct {
	handlers map[string]func(c *Context)
}

func (h *HandleBaseOnMap) key(method string,path string) string {
	return fmt.Sprintf("%s#%s",method,path)
}

func (h *HandleBaseOnMap) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	k := h.key(r.Method,r.URL.Path)
	if handler,ok := h.handlers[k];ok {
		handler(NewContext(w,r))
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("not match any route"))
	}
}

func (h *HandleBaseOnMap) Route(method string,pattern string,handlerFunc func(c *Context))  {
	key := h.key(method,pattern)
	h.handlers[key] = handlerFunc
}

func (s *sdkHttpServer) Route(method string,pattern string,handlerFunc func(c *Context)) {
	s.handler.Route(method,pattern,handlerFunc)
}

func (s *sdkHttpServer) Start(addr string) error {
	fmt.Printf("listening %s\n",addr)
	return http.ListenAndServe(addr,s.handler)
}

// NewServer 新建实例
func NewServer(name string) *sdkHttpServer{
	return &sdkHttpServer{
		Name: name,
		handler: &HandleBaseOnMap{handlers: make(map[string]func(c *Context))},
	}
}

// NewRes 响应格式
func NewRes(code int,msg string,data interface{}) *commonResponse{
	return &commonResponse{
		BizCode: code,
		Msg: msg,
		Data: data,
	}
}

