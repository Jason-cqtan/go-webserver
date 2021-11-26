package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	body,err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w,"read body fail:%v",err)
		return
	}
	err = json.Unmarshal(body,req)
	fmt.Println(err)
	if err != nil {
		log.Printf("error decoding sakura response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("sakura response: %q", body)
		fmt.Fprintf(w,"error decoding sakura response:%v",err)
		return
	}

	// 假设返回正确id
	fmt.Fprintf(w,"id:%d",1)

}


func main() {
	server := &sdkHttpServer{
		name: "openresty",
	}

	server.Route("/signup",SignUp)
	server.Start(":8099")
}
