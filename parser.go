package httpserver

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

type Header map[string]string

type HttpRequest struct {
	Method  string
	Path    string
	Version string
	Headers []Header

	ContentLength int64
	ContentType   string

	// why not a string? the std lib implements it this way, so just a opportunity for me to learn
	// How is this better?
	// if we make this a string then we store the entire request body in Memory while serving a request
	// this is bad because again we will parse the string to create a data structure with same bytes
	// The moment we establish a TCP connection and the clients sends the HTTP request we already have the request in our Memory
	// So O(n) space is already occcupied, now parsing the request should be optmizing memory more and more
	// we can say O(3n) space if we use String, else O(n) when using Reader/Writer

	// Another reason is the optimization the body parser/decoder like json decoder, xml decoder offers which
	// takes in a Reader/Writer interface to perform read the body using *academic* parsing algorithms
	Body io.ReadCloser
}

/*
HTTP REQUEST STRUCTURE:

METHOD path http_version    (Called RequestLine)
KEY1 VAL1
KEY2 VAL2

body
*/
func Parse(conn net.Conn) (*HttpRequest, error) {
	reader := bufio.NewReader(conn)

	method, path, version, err := getRequestLine(reader)
	if err != nil {
		return nil, err
	}
	headers, err := getHeaders(reader)
	if err != nil {
		return nil, err
	}

	// having a body in a request is not mandatory, so lets be careful
	// the fd of our connection is now at the start of the body, as the getHeaders already read the empty line
	// which is the seperator for the body from the header
	contentLength, err := getContentLength(headers)
	if err != nil {
		return nil, err
	}
	contentType := getContentType(headers)

	body := getBody(reader, contentLength)

	request := &HttpRequest{
		Method:        method,
		Path:          path,
		Version:       version,
		Headers:       headers,
		ContentLength: contentLength,
		ContentType:   contentType,

		Body: body,
	}
	return request, nil

}

func getRequestLine(reader *bufio.Reader) (method, path, version string, err error) {
	requestLine, err := reader.ReadString('\n')

	if err != nil {
		return "", "", "", err
	}
	parts := strings.SplitN(requestLine, " ", 3)
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("invalid request line")
	}
	method = parts[0]
	path = parts[1]
	version = strings.TrimRight(parts[2], "\r\n")
	return
}

func getHeaders(reader *bufio.Reader) (headers []Header, err error) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		if line == "\r" || line == "\r\n" {
			break
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid header line")
		}
		fmt.Println(headers)
		headers = append(headers, map[string]string{strings.TrimSpace(parts[0]): strings.TrimSpace(parts[1])})
	}
	return
}

func getContentLength(headers []Header) (int64, error) {
	for _, header := range headers {
		if value, exists := header["Content-Length"]; exists {
			contentLength, err := strconv.Atoi(strings.TrimSpace(value))
			if err != nil {
				return 0, fmt.Errorf("cannot parse content length")
			}
			return int64(contentLength), nil
		}
	}
	return 0, nil
}

func getContentType(headers []Header) string {
	for _, header := range headers {
		if value, exists := header["Content-Type"]; exists {
			return value
		}
	}
	return ""
}

func getBody(reader *bufio.Reader, contentLength int64) io.ReadCloser {
	return io.NopCloser(io.LimitReader(reader, contentLength))
}
