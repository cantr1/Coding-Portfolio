package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strings"
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

func (cfg *apiConfig) resetServerHits() {
	cfg.fileServerHits.Store(0)
}

func main() {
	var apiCfg = apiConfig{}
	const port = ":8080"
	const filepathRoot = "./web/"

	var mux = http.NewServeMux()

	// Health Check
	mux.HandleFunc("GET /api/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // Normal header
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "OK")
	})

	// File Server
	//mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot))))
	handler := http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot)))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handler))

	// Metrics
	mux.HandleFunc("GET /admin/metrics", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8") // Normal header
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, fmt.Sprintf(`<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`, apiCfg.fileServerHits.Load()))
	})

	// Reset
	mux.HandleFunc("POST /admin/reset", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // Normal header
		w.WriteHeader(http.StatusOK)
		apiCfg.resetServerHits()
		io.WriteString(w, "Reset Server hits to 0\n")
	})

	// Validate POST Data
	mux.HandleFunc("POST /api/validate_chirp", func(w http.ResponseWriter, req *http.Request) {
		type parameters struct {
			Body string `json:"body"`
		}

		type returnError struct {
			Error string `json:"error"`
		}

		// type returnValid struct {
		// 	Valid bool `json:"valid"`
		// }

		type returnCleaned struct {
			Body string `json:"cleaned_body"`
		}

		w.Header().Set("Content-Type", "application/json") // All repsonses will be JSON

		decoder := json.NewDecoder(req.Body)
		params := parameters{}
		err := decoder.Decode(&params)
		if err != nil || params.Body == "" { // handle errors and empty POSTs
			log.Printf("Error decoding parameters: %s", err)
			respBody := returnError{
				Error: "Something went wrong",
			}
			dat, err := json.Marshal(respBody)
			if err != nil {
				log.Printf("Error marshalling JSON: %s", err)
			}

			w.WriteHeader(500)
			w.Write(dat)
			return
		}

		chirp_len := len(params.Body)
		if chirp_len > 140 {
			log.Printf("Recieved too long string")
			respBody := returnError{
				Error: "Chirp is too long",
			}
			dat, err := json.Marshal(respBody)
			if err != nil {
				log.Printf("Error marshalling JSON: %s", err)
			}
			w.WriteHeader(400)
			w.Write(dat)
			return
		}

		cleanedString := ""
		bannedWords := []string{"kerfuffle", "sharbert", "fornax"}

		for _, word := range strings.Fields(params.Body) {
			if slices.Contains(bannedWords, strings.ToLower(word)) {
				cleanedString = cleanedString + " ****"
			} else {
				cleanedString = cleanedString + " " + word
			}
		}

		// Clean leading white space
		cleanedString = strings.TrimSpace(cleanedString)

		respBody := returnCleaned{
			Body: cleanedString,
		}

		dat, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
		w.Write(dat)
	})

	// Start server
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
