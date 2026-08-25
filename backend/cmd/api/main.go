package main

import (
	"log"
	"net/http"

	"github.com/mclanxx/kubestream/backend/internal/health"
)

func main() {
	http.HandleFunc("/health", health.Handler)

	log.Println("KubeStream API listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
