package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Get("/", func(c *gee.Context) {
		c.Html(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.Get("/:lang/doc", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.Get("/p/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": "123",
			"password": "sjdkj",
		})
	})

	r.Run(":9999")
}
