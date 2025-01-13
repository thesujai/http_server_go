package httpserver_test

import (
	"fmt"
	"net"
	"testing"
	"time"

	httpserver "github.com/thesujai/http_server_go"
)

func TestServer(t *testing.T) {
	server := httpserver.NewServer()

	server.Handle("GET", "/", func(req *httpserver.HttpRequest, res *httpserver.Response) {
		fmt.Println("GET", req.Path)
		res.Write(200, "Welcome!")
	})

	errCh := make(chan error, 1)

	go func() {
		err := server.ListenAndServe(":8081")
		errCh <- err
	}()

	time.Sleep(100 * time.Millisecond)

	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil {
		t.Fatalf("failed to write: %v", err)
	}
	time.Sleep(100 * time.Millisecond)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatalf("failed to read: %v", err)
	}

	expected := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 8\r\n\r\nWelcome!"
	if string(buf[:n]) != expected {
		t.Errorf("expected response: %q, got: %q", expected, string(buf[:n]))
	}

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("server error: %v", err)
		}
	default:
	}
}
