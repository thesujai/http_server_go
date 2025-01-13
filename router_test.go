package httpserver_test

import (
	"bytes"
	"testing"

	httpserver "github.com/thesujai/http_server_go"
	"github.com/thesujai/http_server_go/mocks"
)

func TestRouter(t *testing.T) {
	router := httpserver.NewRouter()

	router.Handle("GET", "/hello", func(req *httpserver.HttpRequest, res *httpserver.Response) {
		res.Write(200, "Hello, world!")
	})

	req := &httpserver.HttpRequest{
		Method: "GET",
		Path:   "/hello",
	}
	buf := &bytes.Buffer{}
	res := httpserver.NewResponse(&mocks.MockConn{Writer: buf})

	router.Serve(req, res)

	expected := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 13\r\n\r\nHello, world!"
	if buf.String() != expected {
		t.Errorf("expected response: %q, got: %q", expected, buf.String())
	}
}
