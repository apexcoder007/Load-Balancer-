package main

import (
	"fmt"
	"net/http"
)

func main() {
	servers := []Server{
		newSimpleServer("https://www.duckduckgo.com"),
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("https://www.google.com"),
	}

	lb := NewLoadBalancer("8000", servers)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lb.serveProxy(w, r)
	})

	fmt.Printf("Load balancer listening on localhost:%s\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
