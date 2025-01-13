package httpserver_test

import (
	"bytes"
	"testing"

	httpserver "github.com/thesujai/http_server_go"
)

func TestMiddlewareChain(t *testing.T) {
	chain := httpserver.NewMiddlewareChain()

	output := &bytes.Buffer{}
	chain.Use(func(req *httpserver.HttpRequest, res *httpserver.Response, next func()) {
		output.WriteString("Middleware 1 -> ")
		next()
		output.WriteString(" <- Middleware 1")
	})

	chain.Use(func(req *httpserver.HttpRequest, res *httpserver.Response, next func()) {
		output.WriteString("Middleware 2")
		next()
	})

	final := func() {
		output.WriteString("Final Handler")
	}

	chain.Execute(nil, nil, final)

	expected := "Middleware 1 -> Middleware 2Final Handler <- Middleware 1"
	if output.String() != expected {
		t.Errorf("expected output: %q, got: %q", expected, output.String())
	}
}
