package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type RouterGroup struct {
	prefix string
	middlewares []HandlerFunc // support middleware
	parent *RouterGroup   // support nesting
	engine *Engine // all groups share a engine instance
}


type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups
}

func New() *Engine {
	engine := &Engine{router:newRouter()}
	engine.RouterGroup = &RouterGroup{engine:engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix:      group.prefix+prefix,
		parent:      group,
		engine:      engine,
	}
	engine.groups =append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRouter(method, pattern, handler)
}

func (group *RouterGroup) GET (pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST (pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) Run(addr string) (err error) {
	return http.ListenAndServe(addr, group)
}

func (engine *Engine) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}



