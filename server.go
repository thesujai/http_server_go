package httpserver

import (
	"net"
)

type Server struct {
	router *Router
}

func NewServer() *Server {
	return &Server{
		router: NewRouter(),
	}
}

func (s *Server) Handle(method, path string, handler HandlerFunc) {
	s.router.Handle(method, path, handler)
}

func (s *Server) ListenAndServe(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	httpRequest, err := Parse(conn)
	if err != nil {
		return
	}

	res := NewResponse(conn)
	s.router.Serve(httpRequest, res)
}
