package main

import (
	gee "WebGogo/gee"
	"fmt"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "url.path = %q\n", req.URL.Path)
	})

	r.Get("/hellow", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	})
	r.Run(":8888")
}