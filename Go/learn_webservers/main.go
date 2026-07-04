package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
	"sync/atomic"
	"time"

	"example.com/learn_web_servers/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileServerHits atomic.Int32
	dbQueries      database.Queries
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
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
	godotenv.Load() // load env vars
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")

	// open connection to database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Failure to open connection to backend DB")
		return
	}

	var apiCfg = apiConfig{}
	apiCfg.dbQueries = *database.New(db) // attach db queries so handlers can access
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
		if platform != "dev" {
			http.Error(w, "Application not in dev environment", http.StatusForbidden)
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // Normal header
		apiCfg.resetServerHits()
		err := apiCfg.dbQueries.RemoveUsers(req.Context())
		if err != nil {
			http.Error(w, "Failure to erase SQL DB", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Reset Server hits to 0 and erased DB\n")
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

		w.Header().Set("Content-Type", "application/json") // All responses will be JSON

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
				http.Error(w, "Error marshaling JSON string", http.StatusInternalServerError)
			}

			w.WriteHeader(400)
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

	// Create User
	mux.HandleFunc("POST /api/users", func(w http.ResponseWriter, req *http.Request) {
		type parameters struct {
			Email string `json:"email"`
		}

		w.Header().Set("Content-Type", "application/json")

		// decode JSON
		decoder := json.NewDecoder(req.Body)
		params := parameters{}
		err := decoder.Decode(&params)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if params.Email == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}

		dbUser, err := apiCfg.dbQueries.CreateUser(req.Context(), params.Email)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Write dbUser data to our controlled struct
		// to avoid over exposing data via the API
		user := User{
			ID:        dbUser.ID,
			CreatedAt: dbUser.CreatedAt,
			UpdatedAt: dbUser.UpdatedAt,
			Email:     dbUser.Email,
		}

		dat, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Error constructing JSON", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(201)
		w.Write(dat)
	})

	// Start server
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
