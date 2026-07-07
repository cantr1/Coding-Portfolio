package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
	"sync/atomic"
	"time"

	"example.com/learn_web_servers/internal/auth"
	"example.com/learn_web_servers/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileServerHits atomic.Int32
	dbQueries      database.Queries
	tokenSecret    string
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

type Login struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
}

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
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

func validate_chirp(chirp string) bool {
	chirp_len := len(chirp)
	if chirp_len > 140 {
		return false
	}
	return true
}

func clean_chirp(chirp string) string {
	cleanedString := ""
	bannedWords := []string{"kerfuffle", "sharbert", "fornax"}

	for _, word := range strings.Fields(chirp) {
		if slices.Contains(bannedWords, strings.ToLower(word)) {
			cleanedString = cleanedString + " ****"
		} else {
			cleanedString = cleanedString + " " + word
		}
	}

	// Clean leading white space
	cleanedString = strings.TrimSpace(cleanedString)

	return cleanedString
}

func main() {
	godotenv.Load() // load env vars
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	tokenSecret := os.Getenv("JWT")
	tokenDuration := 3600

	// open connection to database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Failure to open connection to backend DB")
		return
	}

	var apiCfg = apiConfig{}
	apiCfg.dbQueries = *database.New(db) // attach db queries so handlers can access
	apiCfg.tokenSecret = tokenSecret
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

	// Get Chirps
	mux.HandleFunc("GET /api/chirps", func(w http.ResponseWriter, req *http.Request) {
		dbChirps, err := apiCfg.dbQueries.GetChirps(req.Context())
		if err != nil {
			http.Error(w, "Unable to retrieve chirps from backend DB", http.StatusInternalServerError)
			return
		}

		processedChirps := []Chirp{}

		for _, chirp := range dbChirps {
			processedChirp := Chirp{
				ID:        chirp.ID,
				CreatedAt: chirp.CreatedAt,
				UpdatedAt: chirp.UpdatedAt,
				Body:      chirp.Body,
				UserID:    chirp.UserID,
			}
			processedChirps = append(processedChirps, processedChirp)
		}

		dat, err := json.Marshal(processedChirps)
		if err != nil {
			http.Error(w, "Unable to marshal chirps to JSON", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
		w.Write(dat)
	})

	// Get chirp by ID
	mux.HandleFunc("GET /api/chirps/{chirpID}", func(w http.ResponseWriter, req *http.Request) {
		chirpIDStr := req.PathValue("chirpID")
		chirpID, err := uuid.Parse(chirpIDStr)
		if err != nil {
			http.Error(w, "Unable to parse ID", http.StatusBadRequest)
			return
		}

		dbChirp, err := apiCfg.dbQueries.GetChirp(req.Context(), chirpID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "ID not in backend DB", http.StatusNotFound)
				return
			}
			http.Error(w, "Unable to retrieve chirps from backend DB", http.StatusInternalServerError)
			return
		}

		processedChirp := Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			Body:      dbChirp.Body,
			UserID:    dbChirp.UserID,
		}

		dat, err := json.Marshal(processedChirp)
		if err != nil {
			http.Error(w, "Unable to marshal chirps to JSON", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
		w.Write(dat)
	})

	// Create a chirp
	mux.HandleFunc("POST /api/chirps", func(w http.ResponseWriter, req *http.Request) {
		type parameters struct {
			Body string `json:"body"`
		}

		w.Header().Set("Content-Type", "application/json") // All responses will be JSON

		decoder := json.NewDecoder(req.Body)
		params := parameters{}
		err = decoder.Decode(&params)
		if err != nil {
			log.Printf("Error decoding parameters: %v", err)
			http.Error(w, "Error decoding parameters", http.StatusBadRequest)
			return
		}

		if params.Body == "" {
			http.Error(w, "Invalid JSON - `body` required", http.StatusBadRequest)
			return
		}

		// Validate token
		passedToken, err := auth.GetBearerToken(req.Header)
		if err != nil {
			http.Error(w, "Auth headers required to post data", http.StatusBadRequest)
			return
		}

		tokenUserUUID, err := auth.ValidateJWT(passedToken, tokenSecret)
		if err != nil {
			log.Printf("Error validating JWT Token: %v", err)
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		valid_chirp := validate_chirp(params.Body)
		if !valid_chirp {
			http.Error(w, "Invalid chirp", http.StatusBadRequest)
			return
		}

		cleanedString := clean_chirp(params.Body)

		// Write chirp to the DB
		chirpParams := database.CreateChirpParams{
			Body:   cleanedString,
			UserID: tokenUserUUID,
		}
		dbChirp, err := apiCfg.dbQueries.CreateChirp(req.Context(), chirpParams)
		if err != nil {
			log.Printf("CreateChirp Error: %v", err)
			http.Error(w, "Error writing chirp to the backend DB", http.StatusInternalServerError)
			return
		}

		chirp := Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			Body:      dbChirp.Body,
			UserID:    dbChirp.UserID,
		}

		dat, err := json.Marshal(chirp)
		if err != nil {
			http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(201)
		w.Write(dat)
	})

	// Create User
	mux.HandleFunc("POST /api/users", func(w http.ResponseWriter, req *http.Request) {
		type parameters struct {
			Email    string `json:"email"`
			Password string `json:"password"`
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

		if params.Password == "" {
			http.Error(w, "Password is required", http.StatusBadRequest)
			return
		}

		// hash password for storage in DB
		hashed_password, err := auth.HashPassword(params.Password)
		if err != nil {
			http.Error(w, "Unable to hash provided password", http.StatusInternalServerError)
			return
		}

		// map parameters to create user struct
		createUserParams := database.CreateUserParams{
			Email:    params.Email,
			Password: hashed_password,
		}

		// Create user
		dbUser, err := apiCfg.dbQueries.CreateUser(req.Context(), createUserParams)
		if err != nil {
			log.Printf("Error: %v", err)
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

	// Update User Info
	mux.HandleFunc("PUT /api/users", func(w http.ResponseWriter, req *http.Request) {
		// Get Token from header
		passedToken, err := auth.GetBearerToken(req.Header)
		if err != nil || passedToken == "" {
			http.Error(w, "Auth Headers with Token Required", http.StatusUnauthorized)
			return
		}

		// Check Token is valid
		dbUserUUID, err := auth.ValidateJWT(passedToken, apiCfg.tokenSecret)
		if err != nil {
			if errors.Is(err, auth.ErrInvalidToken) {
				http.Error(w, "Token invalid", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Error validating JWT", http.StatusUnauthorized) // Unauth for Bootdev
			return
		}

		// Parse request body
		type parameters struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		// decode JSON
		decoder := json.NewDecoder(req.Body)
		params := parameters{}
		err = decoder.Decode(&params)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if params.Email == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}

		if params.Password == "" {
			http.Error(w, "Password is required", http.StatusBadRequest)
			return
		}

		// Hash the PW
		hashedPW, err := auth.HashPassword(params.Password)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		// Update information in the DB
		userInfoParams := database.UpdateUserInfoParams{
			Email:    params.Email,
			Password: hashedPW,
			ID:       dbUserUUID,
		}

		userInfo, err := apiCfg.dbQueries.UpdateUserInfo(req.Context(), userInfoParams)
		if err != nil {
			http.Error(w, "Error writing information to the DB", http.StatusInternalServerError)
			return
		}

		// Parse to struct and marshal to JSON
		userData := User{
			ID:        userInfo.ID,
			Email:     userInfo.Email,
			CreatedAt: userInfo.CreatedAt,
			UpdatedAt: userInfo.UpdatedAt,
		}

		dat, err := json.Marshal(userData)
		if err != nil {
			http.Error(w, "Error marshaling user data to JSON", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
		w.Write(dat)
	})

	// Login endpoint
	mux.HandleFunc("POST /api/login", func(w http.ResponseWriter, req *http.Request) {
		type parameters struct {
			Password string `json:"password"`
			Email    string `json:"email"`
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

		if params.Password == "" {
			http.Error(w, "Password is required", http.StatusBadRequest)
			return
		}

		// Grab User data via email
		userDB, err := apiCfg.dbQueries.QueryUser(req.Context(), params.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "User not found in DB", http.StatusNotFound)
				return
			}
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Check passwords match
		match, err := auth.CheckPasswordHash(params.Password, userDB.Password)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !match {
			http.Error(w, "Incorrect Password", http.StatusUnauthorized)
			return
		}

		// Generate JWT - time.Duration defaults to nanoseconds (that was fun to learn)
		token, err := auth.MakeJWT(userDB.ID, tokenSecret, time.Duration(tokenDuration)*time.Second)
		if err != nil {
			http.Error(w, "Unable to generate JWT", http.StatusInternalServerError)
			return
		}

		// Generate Refresh Token
		refreshToken := auth.MakeRefreshToken()
		// Insert Refresh Token into DB
		refreshTokenParams := database.CreateRefreshTokenParams{
			Token:  refreshToken,
			UserID: userDB.ID,
		}
		_, err = apiCfg.dbQueries.CreateRefreshToken(req.Context(), refreshTokenParams)

		// Successful login - return user data w/o PW
		user := Login{
			ID:           userDB.ID,
			CreatedAt:    userDB.CreatedAt,
			UpdatedAt:    userDB.UpdatedAt,
			Email:        userDB.Email,
			Token:        token,
			RefreshToken: refreshToken,
		}

		dat, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Internal server error - unable to marshal user data", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(dat)
	})

	// Refresh
	mux.HandleFunc("POST /api/refresh", func(w http.ResponseWriter, req *http.Request) {
		// Get Refresh Token from header
		passedToken, err := auth.GetBearerToken(req.Header)
		if err != nil || passedToken == "" {
			http.Error(w, "Auth Headers with Token Required", http.StatusBadRequest)
			return
		}

		// Check Refresh Token Exists
		refreshTokenDB, err := apiCfg.dbQueries.QueryRefreshToken(req.Context(), passedToken)
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		// Check Refresh Token not revoked / expired
		now := time.Now()
		if refreshTokenDB.RevokedAt.Valid {
			http.Error(w, "Token revoked", http.StatusUnauthorized)
			return
		}

		if now.After(refreshTokenDB.ExpiresAt) {
			http.Error(w, "Token expired", http.StatusUnauthorized)
			return
		}

		// Generate New JWT
		token, err := auth.MakeJWT(refreshTokenDB.UserID, tokenSecret, time.Duration(tokenDuration)*time.Second)
		if err != nil {
			http.Error(w, "Unable to generate JWT", http.StatusInternalServerError)
			return
		}

		type RespToken struct {
			Token string `json:"token"`
		}

		respBody := RespToken{
			Token: token,
		}

		dat, err := json.Marshal(respBody)
		if err != nil {
			http.Error(w, "Unable to marshal token into JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dat)
	})

	// Revoke
	mux.HandleFunc("POST /api/revoke", func(w http.ResponseWriter, req *http.Request) {
		// Get Refresh Token from header
		passedToken, err := auth.GetBearerToken(req.Header)
		if err != nil || passedToken == "" {
			http.Error(w, "Auth Headers with Token Required", http.StatusBadRequest)
			return
		}

		// Check Refresh Token Exists
		refreshTokenDB, err := apiCfg.dbQueries.QueryRefreshToken(req.Context(), passedToken)
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		// Check Refresh Token not already revoked
		if refreshTokenDB.RevokedAt.Valid {
			http.Error(w, "Token already revoked", http.StatusBadRequest)
			return
		}

		// Revoke in DB
		err = apiCfg.dbQueries.RevokeRefreshToken(req.Context(), refreshTokenDB.Token)
		if err != nil {
			http.Error(w, "Error revoking token in internal DB", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})

	// Start server
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
