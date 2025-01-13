package httpserver_test

import (
	"bytes"
	"testing"

	httpserver "github.com/thesujai/http_server_go"
	"github.com/thesujai/http_server_go/mocks"
)

func TestWriteHeader(t *testing.T) {
	buf := &bytes.Buffer{}
	conn := &mocks.MockConn{Writer: buf}
	res := httpserver.NewResponse(conn)

	headers := httpserver.Header{"Content-Type": "text/plain"}
	err := res.WriteHeader(200, headers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\n"
	if buf.String() != expected {
		t.Errorf("expected response: %q, got: %q", expected, buf.String())
	}
}

func TestWrite(t *testing.T) {
	buf := &bytes.Buffer{}
	conn := &mocks.MockConn{Writer: buf}
	res := httpserver.NewResponse(conn)

	err := res.Write(200, "Hello, world!")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 13\r\n\r\nHello, world!"
	if buf.String() != expected {
		t.Errorf("expected response: %q, got: %q", expected, buf.String())
	}
}
