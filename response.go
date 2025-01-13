package httpserver

import (
	"fmt"
	"net"
)

type Response struct {
	conn net.Conn
}

func NewResponse(conn net.Conn) *Response {
	return &Response{conn: conn}
}

func (r *Response) WriteHeader(statusCode int, headers Header) error {
	_, err := fmt.Fprintf(r.conn, "HTTP/1.1 %d %s\r\n", statusCode, statusText(statusCode))
	if err != nil {
		return err
	}
	for key, value := range headers {
		_, err := fmt.Fprintf(r.conn, "%s: %s\r\n", key, value)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprint(r.conn, "\r\n")
	return err
}

func (r *Response) Write(statusCode int, body string) error {
	headers := Header{
		"Content-Type":   "text/plain",
		"Content-Length": fmt.Sprintf("%d", len(body)),
	}
	if err := r.WriteHeader(statusCode, headers); err != nil {
		return err
	}
	_, err := fmt.Fprint(r.conn, body)
	return err
}

func statusText(code int) string {
	switch code {
	case 200:
		return "OK"
	case 400:
		return "Bad Request"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	default:
		return "Unknown"
	}
}
