**Project Title: Load Balancer**

## Overview

This project implements a simple reverse-proxy load balancer in Go. It distributes incoming HTTP requests across multiple backend servers using a round-robin algorithm. The goal is to demonstrate core load-balancing concepts and provide a lightweight, extensible example that can be adapted to more advanced scenarios.

---

## What Is a Load Balancer?

A load balancer is a critical component in distributed systems and microservices architectures. It sits between clients and a pool of servers, forwarding incoming requests to healthy servers based on a scheduling strategy. Key benefits include:

* **Scalability**: Distributes traffic to multiple servers to handle high load.
* **High Availability**: Redirects traffic away from unhealthy or overloaded instances.
* **Fault Tolerance**: Automatically retries or reroutes failed requests.

Common algorithms include round-robin, least connections, and IP hash. This example uses the simple round-robin approach.

---

## Project Features

* **Reverse Proxy**: Uses Go’s `httputil.ReverseProxy` to forward client requests.
* **Round-Robin Scheduling**: Evenly distributes requests across backends.
* **Health Checks (Stubbed)**: Interface allows for custom health-check logic.
* **Extensible**: Easily add new balancing algorithms or monitoring hooks.

---

## Prerequisites

* Go 1.16+

---

## Installation & Running

1. **Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/Load-Balancer-.git
   cd Load-Balancer-
   ```

2. **Build and Run**:

   ```bash
   go build -o lb main.go server.go loadbalancer.go 
   ./lb
   ```

3. **Access**:
   Open [http://localhost:8000](http://localhost:8000) in your browser; requests will be forwarded to the configured backends.

---

## Code Structure

```text
├── main.go           # Entry point: sets up servers and starts HTTP listener
├── server.go         # Server interface and simpleServer implementation
├── loadbalancer.go   # LoadBalancer struct: round-robin logic and proxy handling
└── README.md         # Project documentation
```

### Key Components

* **Server Interface** (`Address()`, `isAlive()`, `Serve()`): Abstraction for backend targets.
* **simpleServer**: Wraps `httputil.NewSingleHostReverseProxy`, always marked alive.
* **LoadBalancer**:

  * `getNextAvailableServer()`: Picks next alive server.
  * `serveProxy()`: Logs and forwards requests.

---
