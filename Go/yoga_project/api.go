package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	auth "github.com/cantr1/yoga/internal"
	"github.com/cantr1/yoga/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type apiMetrics struct {
	fileServerHits         atomic.Int32
	userCreationHits       atomic.Int32
	instructorCreationHits atomic.Int32
	sessionCreationHits    atomic.Int32
	classRegistrationHits  atomic.Int32
}

type apiMetricsResponse struct {
	FileServerHits         int32 `json:"file_server_hits"`
	UserCreationHits       int32 `json:"user_creation_hits"`
	InstructorCreationHits int32 `json:"instructor_creation_hits"`
	SessionCreationHits    int32 `json:"session_creation_hits"`
	ClassRegistrationHits  int32 `json:"class_registration_hits"`
}

func (m *apiMetrics) snapshot() apiMetricsResponse {
	return apiMetricsResponse{
		FileServerHits:         m.fileServerHits.Load(),
		UserCreationHits:       m.userCreationHits.Load(),
		InstructorCreationHits: m.instructorCreationHits.Load(),
		SessionCreationHits:    m.sessionCreationHits.Load(),
		ClassRegistrationHits:  m.classRegistrationHits.Load(),
	}
}

type apiCfg struct {
	Port                    string
	FilepathRoot            string
	UserCreationToken       string
	InstructorCreationToken string
	AdminKey                string
	dbQueries               database.Queries
	dbURL                   string
	tokenDuration           int
	tokenSecret             string
	metrics                 apiMetrics
}

type User struct {
	ID        uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Login struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
}

func logRequestError(req *http.Request, message string, err error) {
	log.Printf(
		"method=%s path=%s pattern=%q message=%q error=%v",
		req.Method,
		req.URL.Path,
		req.Pattern,
		message,
		err,
	)
}

func main() {
	// Load environment variables and parse to apiCfg
	godotenv.Load()

	var apiCfg apiCfg
	apiCfg.Port = os.Getenv("PORT")
	apiCfg.FilepathRoot = os.Getenv("FILEPATH_ROOT")
	apiCfg.UserCreationToken = os.Getenv("USER_CREATION_TOKEN")
	apiCfg.AdminKey = os.Getenv("ADMIN_KEY")
	apiCfg.tokenDuration, _ = strconv.Atoi(os.Getenv("TOKEN_DURATION")) // Defaults to 3600 - one hour
	apiCfg.tokenSecret = os.Getenv("TOKEN_SECRET")
	apiCfg.dbURL = os.Getenv("DB_URL")
	apiCfg.metrics = apiMetrics{} // struct to track api usage

	// Open DB Connection
	db, err := sql.Open("postgres", apiCfg.dbURL)
	if err != nil {
		log.Printf("Failure to open connection to backend DB: %v", err)
		return
	}

	// This attribute allows the API to interact with DB methods
	apiCfg.dbQueries = *database.New(db)

	// Create a multiplexer to handle http server
	var mux = http.NewServeMux()

	// Start Server
	server := &http.Server{
		Addr:    apiCfg.Port,
		Handler: mux,
	}

	// === API Endpoint Definitions ===
	// Check server health
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		type StatusOKMSG struct {
			Status string `json:"status"`
		}
		response := StatusOKMSG{
			Status: "online",
		}
		data, err := json.Marshal(response)
		if err != nil {
			logRequestError(req, "error marshaling response to JSON", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	// View metrics - admin only endpoint
	mux.HandleFunc("GET /api/metrics", func(w http.ResponseWriter, req *http.Request) {
		// Check admin key
		token, err := auth.GetBearerToken(*req)
		if err != nil {
			logRequestError(req, "Error parsing token from request", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		if token != apiCfg.AdminKey {
			logRequestError(req, "Attempt to access metrics with incorrect admin key", err)
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		// Request authorized - parse metrics data and return
		data, err := json.Marshal(apiCfg.metrics.snapshot())
		if err != nil {
			logRequestError(req, "error marshaling response to JSON", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	// Create User

	log.Fatal(server.ListenAndServe())
}
