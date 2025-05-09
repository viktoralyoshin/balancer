package healthcheck

import (
	"balancer/pkg/types"
	"context"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type Checker struct {
	serversRef *atomic.Value
	interval   time.Duration
}

func NewChecker(serversRef *atomic.Value, interval time.Duration) *Checker {
	return &Checker{
		serversRef: serversRef,
		interval:   interval,
	}
}

func (c *Checker) Start() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		servers := c.serversRef.Load().([]*types.Server)
		for _, s := range servers {
			go c.checkServer(s)
		}
	}
}

func (c *Checker) checkServer(server *types.Server) {
	log.Printf("Health checking %s", server.URL.String())
	alive := isServerAlive(server)
	setServiceStatus(server, alive)
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
