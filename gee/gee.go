package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(c *Context)
type H map[string]string
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: NewRouter()}
}

func (engin *Engine) Get(path string, h HandlerFunc) {
	engin.router.addRoute("GET", path, h)
}

func (engin *Engine) Post(path string, h HandlerFunc) {
	engin.router.addRoute("POST", path, h)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	engine.router.handle(c)
}
func (engine *Engine) Run(port string) {
	log.Fatal(http.ListenAndServe(port, engine))

}
