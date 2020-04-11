package gee

import (
	"log"
	"time"
)

func Loger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.NEXT()
		log.Printf("[&d] %d in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

