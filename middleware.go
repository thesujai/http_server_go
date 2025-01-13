package httpserver

type MiddlewareFunc func(*HttpRequest, *Response, func())

type MiddlewareChain struct {
	middlewares []MiddlewareFunc
}

func NewMiddlewareChain() *MiddlewareChain {
	return &MiddlewareChain{}
}

func (c *MiddlewareChain) Use(middleware MiddlewareFunc) {
	c.middlewares = append(c.middlewares, middleware)
}

func (c *MiddlewareChain) Execute(req *HttpRequest, res *Response, final func()) {
	executeMiddleware(0, c.middlewares, req, res, final)
}

func executeMiddleware(index int, middlewares []MiddlewareFunc, req *HttpRequest, res *Response, final func()) {
	if index < len(middlewares) {
		middlewares[index](req, res, func() {
			executeMiddleware(index+1, middlewares, req, res, final)
		})
	} else {
		final()
	}
}
