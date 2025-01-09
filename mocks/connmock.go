package mocks

import (
	"io"
	"net"
	"time"
)

type MockConn struct {
	Reader io.Reader
	Writer io.Writer
}

func (m *MockConn) Read(b []byte) (n int, err error) {
	return m.Reader.Read(b)
}

func (m *MockConn) Write(b []byte) (n int, err error) {
	return m.Writer.Write(b)
}

func (m *MockConn) Close() error {
	return nil
}

func (m *MockConn) LocalAddr() net.Addr {
	return &net.TCPAddr{}
}

func (m *MockConn) RemoteAddr() net.Addr {
	return &net.TCPAddr{}
}

func (m *MockConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetWriteDeadline(t time.Time) error {
	return nil
}
