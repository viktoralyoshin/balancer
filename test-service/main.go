package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")
	HOST, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Response from %s:%s\n", HOST, PORT)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		return
	}
}
