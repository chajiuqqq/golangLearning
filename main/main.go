package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func Logger() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s : %v \n", c.StatusCode, c.Path, time.Since(t))
	}
}

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Fail(500, "server error\n")
		log.Printf("[%d] %s : %v for v2\n", c.StatusCode, c.Path, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.Use(Logger())
	r.GET("/", func(c *gee.Context) {
		c.Html(200, "<h2>/</h2>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	v2.GET("/", func(c *gee.Context) {
		c.Html(http.StatusOK, "v2: <h1>Hello Gee</h1>")
	})
	v2.GET("/:lang/doc", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "v2: hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	v2.GET("/p/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": "123",
			"password": "sjdkj",
		})
	})

	r.Run(":9999")
}
