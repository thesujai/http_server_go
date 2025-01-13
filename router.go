package httpserver

import (
	"fmt"
)

type Router struct {
	routes map[string]HandlerFunc
}

type HandlerFunc func(*HttpRequest, *Response)

func NewRouter() *Router {
	return &Router{routes: make(map[string]HandlerFunc)}
}

func (r *Router) Handle(method, path string, handler HandlerFunc) {
	key := fmt.Sprintf("%s:%s", method, path)
	r.routes[key] = handler
}

func (r *Router) Serve(req *HttpRequest, res *Response) {
	key := fmt.Sprintf("%s:%s", req.Method, req.Path)
	if handler, exists := r.routes[key]; exists {
		handler(req, res)
		return
	}
	res.Write(404, "Not Found")
}
