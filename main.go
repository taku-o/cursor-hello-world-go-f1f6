package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// newServeMux wires up the HTTP routes for the hello-world service.
func newServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHello)
	mux.HandleFunc("/healthz", handleHealth)
	return mux
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, "hello project")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Printf("starting hello-world server on %s", addr)
	if err := http.ListenAndServe(addr, newServeMux()); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
