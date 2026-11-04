package main

import (
	"fmt"
	"log"
	"net/http"
)

func newHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health/live", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	mux.HandleFunc("GET /health/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "READY")
	})

	return mux
}

func main() {
	log.Println("Company Service running on port 8080")

	if err := http.ListenAndServe(":8080", newHandler()); err != nil {
		log.Fatal(err)
	}
}
