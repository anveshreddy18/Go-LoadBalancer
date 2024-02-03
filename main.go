package main

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

// Define Backend struct
type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

// Define ServerList struct
type ServerList struct {
	backends []*Backend
	current  uint64
}

func main() {

	// Create a list of Backend services without the Alive flag

	// Create the serverList struct.

	// Then run the go routine which periodically checks for the health status of the backend services
	// Goroutine constantly periodically flags any unhealthy backend services

	// Now for every request that comes, call the loadbalancer func

}
