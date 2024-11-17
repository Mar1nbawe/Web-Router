package router

import (
	"context"
	"net/http"
)

var MessageCodes = [...]string{"Inactive", "In progress", "Available"}

// ~~~~~ Router ~~~~~ //

type Router struct {
	routes         []RouteEntry
	middleware     []Middleware
	allowedMethods map[string][]string
}

func (rtr *Router) Use(mw Middleware) {
	rtr.middleware = append(rtr.middleware, mw)
}

func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc, code int) {

	for _, mw := range rtr.middleware {
		handlerFunc = mw(handlerFunc)
	}

	e := RouteEntry{
		Method: method,
		Path:   path,
		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "code", code)
			handlerFunc(w, r.WithContext(ctx))
		},
	}
	rtr.routes = append(rtr.routes, e)

	if rtr.allowedMethods == nil {
		rtr.allowedMethods = make(map[string][]string)
	}
	rtr.allowedMethods[path] = append(rtr.allowedMethods[path], method)
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var handler http.HandlerFunc
	var methodNotAllowed bool

	for _, e := range rtr.routes {

		if e.Match(r) {
			handler = e.HandlerFunc
			break
		}
		if e.Path == r.URL.Path {
			methodNotAllowed = true
		}
	}

	if handler == nil {
		if methodNotAllowed {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		} else {
			handler = http.NotFound
		}
		return
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
