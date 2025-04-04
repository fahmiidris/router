package router

import (
	"net/http"
)

type Group struct {
	r *Router
	p string
}

func newGroup(r *Router, path string) *Group {
	if path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}

	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	return &Group{
		r: r,
		p: path,
	}
}

func (group *Group) Group(path string) *Group {
	return newGroup(group.r, group.subPath(path))
}

func (group *Group) Handle(method, path string, handle Handle) {
	group.r.Handle(method, group.subPath(path), handle)
}

func (group *Group) Handler(method, path string, handler http.Handler) {
	group.r.Handler(method, group.subPath(path), handler)
}

func (group *Group) HandlerFunc(method, path string, handler http.HandlerFunc) {
	group.r.HandlerFunc(method, group.subPath(path), handler)
}

func (group *Group) Use(middleware Middleware) {
	group.r.Use(middleware)
}

func (group *Group) GET(path string, handle Handle) {
	group.Handle("GET", path, handle)
}

func (group *Group) HEAD(path string, handle Handle) {
	group.Handle("HEAD", path, handle)
}

func (group *Group) OPTIONS(path string, handle Handle) {
	group.Handle("OPTIONS", path, handle)
}

func (group *Group) POST(path string, handle Handle) {
	group.Handle("POST", path, handle)
}

func (group *Group) PUT(path string, handle Handle) {
	group.Handle("PUT", path, handle)
}

func (group *Group) PATCH(path string, handle Handle) {
	group.Handle("PATCH", path, handle)
}

func (group *Group) DELETE(path string, handle Handle) {
	group.Handle("DELETE", path, handle)
}

func (group *Group) subPath(path string) string {
	if path[0] != '/' {
		panic("path must start with a '/'")
	}

	return group.p + path
}