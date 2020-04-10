package gee

import "net/http"

type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Req *http.Request
	Path string
	Method string
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:     w,
		Req:        req,
		Path:       req.URL.Path,
		Method:     req.Method,
	}
}
