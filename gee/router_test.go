package gee

import (
	"fmt"
	"testing"
)

func (r *router) addTestRoute() {
	r.GET("/:lang/doc", func(c *Context) { fmt.Println("/:lang/doc handler") })
	r.GET("/:lang/help", func(c *Context) { fmt.Println("/:lang/help handler") })
	r.GET("/about", func(c *Context) { fmt.Println("/about handler") })
	r.GET("/my/:name", func(c *Context) { fmt.Println("/my/:name handler") })
	r.GET("/p/*filename", func(c *Context) { fmt.Println("/p/*filename handler") })
}
func TestGetHandler(t *testing.T) {
	r := NewRouter()
	r.addTestRoute()
	if n, _ := r.GetRoute("GET", "/en/doc"); n.pattern != "/:lang/doc" {
		t.Fatal("/en/doc parase error")
	}

	if n, _ := r.GetRoute("GET", "/cn/help"); n.pattern != "/:lang/help" {
		t.Fatal("/cn/help parase error")
	}

	if n, _ := r.GetRoute("GET", "/about"); n.pattern != "/about" {
		t.Fatal("/about parase error")
	}

	if n, _ := r.GetRoute("GET", "/my/jack"); n.pattern != "/my/:name" {
		t.Fatal("/my/jack parase error")
	}

	if n, _ := r.GetRoute("GET", "/p/1.jpg"); n.pattern != "/p/*filename" {
		t.Fatal("/p/1.jpg parase error")
	}

	if n, _ := r.GetRoute("GET", "/p/static/1.jpg"); n.pattern != "/p/*filename" {
		t.Fatal("/p/static/1.jpg parase error")
	}

	//==
	_, h := r.GetRoute("GET", "/en/doc")
	if h == nil {
		t.Fatal("handler error")
	}
	h(nil)

}
