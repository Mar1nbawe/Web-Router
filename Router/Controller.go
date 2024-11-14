package router

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

// ~~~~~ Router ~~~~~ //

type Router struct {
	routes     []RouteEntry
	middleware []Middleware
}

func (rtr *Router) Use(mw Middleware) {
	rtr.middleware = append(rtr.middleware, mw)
}

func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc) {

	for _, mw := range rtr.middleware {
		handlerFunc = mw(handlerFunc)
	}

	e := RouteEntry{
		Method:      method,
		Path:        path,
		HandlerFunc: handlerFunc,
	}
	rtr.routes = append(rtr.routes, e)
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var handler http.HandlerFunc

	for _, e := range rtr.routes {

		if e.Match(r) {
			handler = e.HandlerFunc
			break
		}
	}
	if handler == nil {
		handler = http.NotFound
	}

	for _, mw := range rtr.middleware {
		handler = mw(handler)
	}

	handler.ServeHTTP(w, r)
}

// ~~~~~ RouteEntry ~~~~~ //

type RouteEntry struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

func (ent *RouteEntry) Match(r *http.Request) bool {
	println(r.Method, r.URL.Path)
	if r.Method != ent.Method {
		return false // Method mismatch
	}

	if r.URL.Path != ent.Path {
		return false // Path mismatch
	}

	return true
}
