package proxy

import (
	"balancer/pkg/types"
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"sync/atomic"
	"time"
)

type Balancer struct {
	servers *atomic.Value
	index   uint32
}

func NewBalancer(servers *atomic.Value) *Balancer {
	return &Balancer{servers: servers}
}

func (b *Balancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	servers := b.servers.Load().([]*types.Server)
	available := b.getAvailableServers(servers)

	if len(available) == 0 {
		http.Error(w, "No available servers", http.StatusServiceUnavailable)
		return
	}

	idx := atomic.AddUint32(&b.index, 1) % uint32(len(available))
	server := available[idx]

	if isServerAlive(server) {
		setServiceStatus(server, true)
		proxy := httputil.NewSingleHostReverseProxy(server.URL)
		proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
			http.Error(w, e.Error(), http.StatusServiceUnavailable)
		}
		proxy.ServeHTTP(w, r)
		return
	}

	setServiceStatus(server, false)
	http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
}

func (b *Balancer) getAvailableServers(servers []*types.Server) []*types.Server {
	var available []*types.Server
	for _, s := range servers {
		if atomic.LoadInt32(&s.Available) == 1 {
			available = append(available, s)
		}
	}
	return available
}

func isServerAlive(server *types.Server) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", server.URL.String()+"/health", nil)
	if err != nil {
		return false
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		if resp != nil {
			_ = resp.Body.Close()
		}
		return false
	}

	_ = resp.Body.Close()
	return true
}

func setServiceStatus(server *types.Server, alive bool) {
	prev := atomic.LoadInt32(&server.Available)
	newValue := int32(0)

	if alive {
		newValue = 1
	}

	if prev != newValue {
		if alive {
			log.Printf("Server %s is available", server.URL.String())
		} else {
			log.Printf("Server %s is unavailable", server.URL.String())
		}
		atomic.StoreInt32(&server.Available, newValue)
	}
}
