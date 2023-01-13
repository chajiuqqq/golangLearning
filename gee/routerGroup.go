package gee

import (
	"net/http"
	"path"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	engine      *Engine
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	e := g.engine
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		engine: e,
	}
	e.groups = append(e.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRoute(method string, pattern string, h HandlerFunc) {
	newPattern := g.prefix + pattern
	g.engine.router.addRoute(method, newPattern, h)
}

func (g *RouterGroup) GET(pattern string, h HandlerFunc) {
	g.addRoute("GET", pattern, h)
}

func (g *RouterGroup) POST(pattern string, h HandlerFunc) {
	g.addRoute("POST", pattern, h)
}
func (g *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {

	absolutePath := path.Join(g.prefix, relativePath)
	fileserver := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		filename := c.Param("filepath")
		if _, err := fs.Open(filename); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(200)
		fileserver.ServeHTTP(c.w, c.r)
	}

}

func (g *RouterGroup) Static(relativePath string, root string) {
	handler := g.createStaticHandler(relativePath, http.Dir(root))
	pattern := path.Join(relativePath, "/*filepath")
	g.GET(pattern, handler)
}
