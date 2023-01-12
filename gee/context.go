package gee

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Context struct {
	w          http.ResponseWriter
	r          *http.Request
	Method     string
	Path       string
	StatusCode int
	handlers   []HandlerFunc
	index      int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{w: w, r: r, Method: r.Method, Path: r.URL.Path, index: -1, handlers: make([]HandlerFunc, 0)}
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.w.WriteHeader(code)
}

func (c *Context) Html(code int, content string) {
	c.Status(code)
	c.w.Header().Set("Content-Type", "text/html")
	c.w.Write([]byte(content))
}
func (c *Context) JSON(code int, item interface{}) {
	c.Status(code)
	encoder := json.NewEncoder(c.w)
	if err := encoder.Encode(item); err != nil {
		http.Error(c.w, err.Error(), 500)
	}
}
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.w.Write(data)
}
func (c *Context) String(code int, format string, values ...interface{}) {
	c.Status(code)
	c.w.Header().Set("Content-Type", "text/plain")
	c.w.Write([]byte(fmt.Sprintf(format, values...)))
}
func (c *Context) Query(key string) string {
	return c.r.URL.Query().Get(key)
}
func (c *Context) PostForm(key string) string {
	return c.r.FormValue(key)
}
func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, msg string) {
	c.Status(code)
	log.Println("error, " + msg)
	c.Html(code, msg)
}
