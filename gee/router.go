package gee

import (
	"fmt"
	"log"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}
func (r *router) addRoute(m string, path string, h HandlerFunc) {
	key := m + "-" + path
	r.handlers[key] = h
	if r.roots[m] == nil {
		r.roots[m] = NewNode()
	}
	r.roots[m].insert(path, patternSplit(path))
}

func (r *router) GET(path string, h HandlerFunc) {
	r.addRoute("GET", path, h)
}

func (r *router) POST(path string, h HandlerFunc) {
	r.addRoute("POST", path, h)
}

func (r *router) GetRoute(m string, path string) (*node, HandlerFunc) {
	node := r.roots[m].search(patternSplit(path), 0)
	if node != nil {
		return node, r.handlers[m+"-"+node.pattern]
	}
	return nil, nil
}

func (r *router) handle(c *Context) {

	key := c.Method + "-" + c.Path
	if _, handler := r.GetRoute(c.Method, c.Path); handler != nil {
		handler(c)
	} else {
		log.Println("404 not found: " + key)
		fmt.Fprintf(c.w, "404 not found:%s %s", c.Method, c.Path)
	}
}
