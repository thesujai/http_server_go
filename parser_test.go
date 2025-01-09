package httpserver_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"testing"

	httpserver "github.com/thesujai/http_server_go"
	"github.com/thesujai/http_server_go/mocks"
)

func TestParse(t *testing.T) {
	testData := []struct {
		name                string
		requestData         string
		expectedHttpRequest *httpserver.HttpRequest
		expectError         bool
		err                 error
	}{
		{
			name:        "Valid GET Request",
			requestData: "GET /hello HTTP/1.1\r\nContent-Length: 11\r\nContent-Type: text/plain\r\n\r\nHello World",
			expectedHttpRequest: &httpserver.HttpRequest{
				Method:  "GET",
				Path:    "/hello",
				Version: "HTTP/1.1",
				Headers: []httpserver.Header{
					{"Content-Length": "11"},
					{"Content-Type": "text/plain"},
				},
				ContentLength: 11,
				ContentType:   "text/plain",
				Body:          io.NopCloser(io.LimitReader(bytes.NewReader([]byte("Hello World")), 11)),
			},
			expectError: false,
		},
		{
			name:        "Valid POST Request",
			requestData: "POST /login HTTP/1.1\r\nContent-Length: 45\r\nContent-Type: application/json\r\n\r\n{\"username\": \"johndoe\", \"password\": \"secret\"}",
			expectedHttpRequest: &httpserver.HttpRequest{
				Method:  "POST",
				Path:    "/login",
				Version: "HTTP/1.1",
				Headers: []httpserver.Header{
					{"Content-Length": "45"},
					{"Content-Type": "application/json"},
				},
				ContentLength: 45,
				ContentType:   "application/json",
				Body:          io.NopCloser(io.LimitReader(bytes.NewReader([]byte(`{"username": "johndoe", "password": "secret"}`)), 45)),
			},
			expectError: false,
		},
		{
			name:                "Invalid request line",
			requestData:         "GET /hello\r\nContent-Length: 11\r\nContent-Type: text/plain\r\n\r\nHello World",
			expectedHttpRequest: nil,
			expectError:         true,
			err:                 fmt.Errorf("invalid request line"),
		},
		{
			name:        "Missing Body",
			requestData: "GET /hello HTTP/1.1\r\nContent-Length: 0\r\n\r\n",
			expectedHttpRequest: &httpserver.HttpRequest{
				Method:  "GET",
				Path:    "/hello",
				Version: "HTTP/1.1",
				Headers: []httpserver.Header{
					{"Content-Length": "0"},
				},
				ContentLength: 0,
				ContentType:   "",
				Body:          io.NopCloser(bytes.NewReader(nil)),
			},
		},
	}

	for _, tc := range testData {
		t.Run(tc.name, func(t *testing.T) {
			mockConn := mocks.MockConn{
				Reader: bytes.NewBuffer([]byte(tc.requestData)),
				Writer: &bytes.Buffer{},
			}
			req, err := httpserver.Parse(&mockConn)

			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				if err.Error() != tc.err.Error() {
					t.Fatalf("expected error '%v' but got '%v'", tc.err, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if req.Method != tc.expectedHttpRequest.Method {
				t.Errorf("expected method '%s', got '%s'", tc.expectedHttpRequest.Method, req.Method)
			}
			if req.Path != tc.expectedHttpRequest.Path {
				t.Errorf("expected path '%s', got '%s'", tc.expectedHttpRequest.Path, req.Path)
			}
			if req.Version != tc.expectedHttpRequest.Version {
				t.Errorf("expected version '%s', got '%s'", tc.expectedHttpRequest.Version, req.Version)
			}
			if req.ContentLength != tc.expectedHttpRequest.ContentLength {
				t.Errorf("expected content length '%d', got '%d'", tc.expectedHttpRequest.ContentLength, req.ContentLength)
			}
			if req.ContentType != tc.expectedHttpRequest.ContentType {
				t.Errorf("expected content type '%s', got '%s'", tc.expectedHttpRequest.ContentType, req.ContentType)
			}
			for i, expectedHeader := range tc.expectedHttpRequest.Headers {
				if i >= len(req.Headers) {
					t.Errorf("missing header at index %d: %v", i, expectedHeader)
					continue
				}
				for k, v := range expectedHeader {
					if req.Headers[i][k] != v {
						t.Errorf("expected header '%s: %s', got '%s'", k, v, req.Headers[i][k])
					}
				}
			}
			if tc.expectedHttpRequest.Body != nil {
				expectedBody, _ := ioutil.ReadAll(tc.expectedHttpRequest.Body)
				actualBody, _ := ioutil.ReadAll(req.Body)

				if !bytes.Equal(expectedBody, actualBody) {
					t.Errorf("expected body '%s', got '%s'", string(expectedBody), string(actualBody))
				}
			}

		})
	}

}
