package main

import (
	"balancer/internal/balancer"
	"log"
	"net/http"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")

	configPath := "config.json"

	lb, err := balancer.NewLoadBalancer(configPath)
	if err != nil {
		log.Fatalf("Failed to create load balancer: %v", err)
	}

	log.Printf("Listening on port %s", PORT)
	log.Println("Send SIGHUP to reload config.")
	log.Fatal(http.ListenAndServe(":"+PORT, lb.Handler()))
}
