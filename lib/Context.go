package lib

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func NewContext(w http.ResponseWriter,r *http.Request) *Context {
	return &Context{
		W:w,
		R:r,
	}
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