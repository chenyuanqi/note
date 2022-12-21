package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	// 中间件不仅作用在处理流程前，也可以作用在处理流程后
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
