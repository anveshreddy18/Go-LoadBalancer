package main

import (
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
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

func (b *Backend) setAlive(flag bool) {
	b.mux.Lock()
	b.Alive = flag
	b.mux.Unlock()
}

func (b *Backend) isAlive() (alive bool) {
	b.mux.RLock()
	alive = b.Alive
	b.mux.RUnlock()
	return alive
}

func (s *ServerList) MarkBackendStatus(url *url.URL, flag bool) {
	for _, b := range s.backends {
		if b.URL.String() == url.String() {
			b.setAlive(flag)
			return
		}
	}
}

func (s *ServerList) AddBackend(b *Backend) {
	s.backends = append(s.backends, b)
}

// NextIndex returns the next cyclic index to the s.current.
// The function exists to provide atomicity in the addition operation
func (s *ServerList) NextIndex() int {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
}

func (s *ServerList) GetNextPeer() *Backend {
	nextInd := s.NextIndex()
	// Now loop in a cyclic way to find the next backend service that is alive
	for i := nextInd; i < (len(s.backends) + nextInd); i++ {
		curInd := i % len(s.backends)
		if s.backends[curInd].isAlive() {
			// Store this index in the s.current & return this backend
			atomic.StoreUint64(&s.current, uint64(curInd))
			return s.backends[curInd]
		}
	}
	return nil
}

func main() {

	// Create a list of Backend services without the Alive flag
	// url1, _ := url.Parse("http://localhost:1")
	// url2, _ := url.Parse("http://localhost:2")
	// url3, _ := url.Parse("http://localhost:3")
	// url4, _ := url.Parse("http://localhost:4")

	// backendList := []Backend{{URL: url1}}

	// Create the serverList struct.

	// Then run the go routine which periodically checks for the health status of the backend services
	// Goroutine constantly periodically flags any unhealthy backend services

	// Now for every request that comes, call the loadbalancer func

}
