package gee

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)
type H map[string]string
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New() *Engine {
	e := &Engine{router: NewRouter()}
	e.RouterGroup = &RouterGroup{engine: e}
	e.groups = []*RouterGroup{e.RouterGroup}
	return e
}

func (engin *Engine) GET(path string, h HandlerFunc) {
	engin.router.addRoute("GET", path, h)
}

func (engin *Engine) POST(path string, h HandlerFunc) {
	engin.router.addRoute("POST", path, h)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	//把适配这个path的中间件加进去
	for _, group := range engine.groups {
		if strings.HasPrefix(c.Path, group.prefix) {
			c.handlers = append(c.handlers, group.middlewares...)
		}
	}
	engine.router.handle(c)
}
func (engine *Engine) Run(port string) {
	log.Fatal(http.ListenAndServe(port, engine))

}

// 使用中间件
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
