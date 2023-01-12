package gee

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
