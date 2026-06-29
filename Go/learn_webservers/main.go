package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileServerHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) resetServerHits()  {
	cfg.fileServerHits.Store(0)	
}

func main() {
	var apiCfg = apiConfig{}
	const port = ":8080"
	const filepathRoot = "./web/"

	var mux = http.NewServeMux()

	// Health Check
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // Normal header
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "OK")
	})

	// File Server
	//mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot))))
	handler := http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot)))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handler))

	// Metrics
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // Normal header
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, fmt.Sprintf("Hits: %v\n", apiCfg.fileServerHits.Load()))
	})
	mux.HandleFunc("/reset", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // Normal header
		w.WriteHeader(http.StatusOK)
		apiCfg.resetServerHits()
		io.WriteString(w, "Reset Server hits to 0\n")
	})

	// Start server
	server := &http.Server{
		Addr: port,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}