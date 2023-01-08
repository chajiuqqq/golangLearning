package gee

import (
	"fmt"
	"log"
)

type router struct {
	handlers map[string]HandlerFunc
}

func (r *router) Route(m string, path string, h HandlerFunc) {
	key := m + "-" + path
	r.handlers[key] = h
}
func NewRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if h, ok := r.handlers[key]; ok {
		log.Println(key)
		h(c)
	} else {
		log.Println("404 not found: " + key)
		fmt.Fprintf(c.w, "404 not found:%s %s", c.Method, c.Path)
	}
}
