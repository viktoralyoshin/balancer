package balancer

import (
	"balancer/internal/config"
	"balancer/internal/healthcheck"
	"balancer/internal/proxy"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

type LoadBalancer struct {
	servers atomic.Value
	proxy   *proxy.Balancer
}

func NewLoadBalancer(configPath string) (*LoadBalancer, error) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	servers, err := config.ParseServers(cfg.Servers)
	if err != nil {
		return nil, err
	}

	lb := &LoadBalancer{
		servers: atomic.Value{},
	}

	lb.proxy = proxy.NewBalancer(&lb.servers)

	lb.servers.Store(servers)

	checker := healthcheck.NewChecker(&lb.servers, 10*time.Second)
	go checker.Start()
	lb.setupSignalHandler(configPath)

	return lb, nil
}

func (lb *LoadBalancer) setupSignalHandler(configPath string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	go func() {
		for range c {
			log.Println("Received SIGHUP, reloading config")

			cfg, err := config.LoadConfig(configPath)
			if err != nil {
				log.Printf("Config reload failed: %v", err)
				continue
			}

			newServers, err := config.ParseServers(cfg.Servers)
			if err != nil {
				log.Printf("Config reload failed: %v", err)
				continue
			}

			lb.servers.Store(newServers)

			log.Printf("Config successfully reloaded: %d servers", len(newServers))
		}
	}()
}

func (lb *LoadBalancer) Handler() http.Handler {
	return lb.proxy
}
