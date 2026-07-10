package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	auth "github.com/cantr1/internal"
	"github.com/cantr1/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

// --- Begin Struct Definitions
type apiCfg struct {
	Port         string
	FilepathRoot string
	dbQueries    database.Queries
}

type User struct {
	ID        uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// --- End Struct Definitions

func main() {
	// Load Env Vars
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")

	// open connection to database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Failure to open connection to backend DB")
		return
	}

	// Construct apiCfg for interfacing with env vars and DB
	var apiCfg apiCfg
	apiCfg.Port = os.Getenv("PORT")
	apiCfg.FilepathRoot = os.Getenv("FILEPATH_ROOT")
	apiCfg.dbQueries = *database.New(db)

	// Create a multiplexer
	var mux = http.NewServeMux()

	// --- Begin API Endpoint Definitions

	// Check health of API
	mux.HandleFunc("GET /api/healthy", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		type statusOKMsg struct {
			Status string `json:"string"`
		}
		response := statusOKMsg{
			Status: "online/healthy",
		}
		data, err := json.Marshal(response)
		if err != nil {
			log.Printf("Error marshaling to JSON: %v", err)
			http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	// Create User in DB
	mux.HandleFunc("POST /api/users", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Parse parameters from request
		type parameters struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		decoder := json.NewDecoder(req.Body)
		params := parameters{}
		err := decoder.Decode(&params)
		if err != nil {
			log.Printf("Error decoding parameters: %v", err)
			http.Error(w, "Error decoding parameters", http.StatusBadRequest)
			return
		}

		// Check required parameters exist
		if params.Email == "" {
			http.Error(w, "Invalid JSON - `email` required", http.StatusBadRequest)
			return
		}

		if params.Password == "" {
			http.Error(w, "Invalid JSON - `password` required", http.StatusBadRequest)
			return
		}

		// Hash the user password for the DB
		hashedPW, err := auth.HashPassword(params.Password)
		if err != nil {
			log.Printf("Error hashing PW: %v", err)
			http.Error(w, "Error hashing user password", http.StatusInternalServerError)
			return
		}

		// Parse info into database struct
		dbParams := database.CreateUserParams{
			Email:    params.Email,
			Password: hashedPW,
		}

		// Create user in the DB
		userDB, err := apiCfg.dbQueries.CreateUser(req.Context(), dbParams)

		// Parse returned SQL data into user struct to return in response body
		user := User{
			ID: userDB.ID,
			CreatedAt: userDB.CreatedAt,
			UpdatedAt: userDB.UpdatedAt,
			Email: userDB.Email,
		}

		// Marshal to JSON and return
		data, err := json.Marshal(&user)
		if err != nil {
			log.Printf("Error marshaling JSON into user struct: %v", err)
			http.Error(w, "Error marshaling user data to struct", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	})

	// --- End API Endpoint Definitions

	// Start Server
	server := &http.Server{
		Addr:    apiCfg.Port,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
