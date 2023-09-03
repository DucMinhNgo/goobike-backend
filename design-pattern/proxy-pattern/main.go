package main

import (
	"time"
)

type Server struct {
	fqdn string
	port uint32
	timeout time.Duration
	maxConn uint32
}

type optionalFunc func(*Server)

func NewServer(options ...optionalFunc) *Server {
	svr := &Server{}
	for _, opt := range options {
		opt(svr)
	}
	return svr
}

func WithHost(fqdn string) optionalFunc {
	return func(s *Server) {
    	s.fqdn = fqdn
  	}
}

func WithPort(port uint32) optionalFunc {
	return func(s *Server) {
    	s.port = port
  	}
}

func WithTimeout(timeout time.Duration) optionalFunc {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn uint32) optionalFunc {
	return func(s *Server) {
    	s.maxConn = maxConn
  	}
}

func main() {
	
}