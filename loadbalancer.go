package main

import (
	"fmt"
	"net/http"
)

// LoadBalancer holds state for round-robin distribution.
type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

// NewLoadBalancer initializes the balancer.
func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

// getNextAvailableServer returns the next healthy Server.
func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.isAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

// serveProxy picks a server, logs, and forwards the request.
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("Forwarding request to %q\n", targetServer.Address())
	targetServer.Serve(rw, req)
}
