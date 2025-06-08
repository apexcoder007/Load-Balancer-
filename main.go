package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	isAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

type simpleServer struct {
	addr  string
	proxy httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	serverURL, err := url.Parse(addr)
	handleError(err)

	return &simpleServer{
		addr:  addr,
		proxy: *httputil.NewSingleHostReverseProxy(serverURL),
	}
}

type LoadBalancer struct {
	port           string
	rounRobinCount int
	Servers        []Server
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
