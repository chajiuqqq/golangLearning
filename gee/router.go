package gee

import (
	"fmt"
	"log"
	"strings"
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
	//去除最后的/
	// if strings.HasSuffix(path, "/") {
	// 	path = path[:len(path)-1]
	// }
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

func (r *router) GetRoute(m string, path string) (*node, HandlerFunc, map[string]string) {
	searchParts := patternSplit(path)
	node := r.roots[m].search(patternSplit(path), 0)
	params := make(map[string]string)
	if node != nil {
		parts := patternSplit(node.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
			}
		}
		return node, r.handlers[m+"-"+node.pattern], params
	}
	return nil, nil, nil
}

func (r *router) handle(c *Context) {

	key := c.Method + "-" + c.Path
	if _, handler, params := r.GetRoute(c.Method, c.Path); handler != nil {
		c.handlers = append(c.handlers, handler)
		c.params = params
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			log.Println("404 not found: " + key)
			c.Html(404, fmt.Sprintf("404 not found:%s %s", c.Method, c.Path))
		})
	}
	c.Next()
}
