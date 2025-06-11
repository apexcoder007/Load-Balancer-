package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// Server is a target that can handle proxied requests.
type Server interface {
	Address() string
	isAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

// simpleServer wraps Go's ReverseProxy.
type simpleServer struct {
	addr  string
	proxy httputil.ReverseProxy
}

// newSimpleServer constructs a simpleServer pointing to addr.
func newSimpleServer(addr string) *simpleServer {
	serverURL, err := url.Parse(addr)
	handleError(err)
	return &simpleServer{
		addr:  addr,
		proxy: *httputil.NewSingleHostReverseProxy(serverURL),
	}
}

// Address returns the backend URL.
func (s *simpleServer) Address() string { return s.addr }

// isAlive stub—always healthy for now.
func (s *simpleServer) isAlive() bool { return true }

// Serve uses the embedded ReverseProxy.
func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

// handleError is a small helper to exit on fatal errors.
func handleError(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
