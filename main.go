package main

import (
	"net/http"

	"WebGogo/gogo"
)

func main() {
	r := gogo.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}