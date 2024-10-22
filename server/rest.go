package server

import (
	"log"
	"net/http"
)

func StartRESTServer(port string) {
	log.Printf("REST server is running on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("failed to start REST server: %v", err)
	}
}
