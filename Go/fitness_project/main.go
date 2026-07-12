package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	auth "github.com/cantr1/internal"
	"github.com/cantr1/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// --- Begin Struct Definitions
type apiCfg struct {
	Port              string
	FilepathRoot      string
	UserCreationToken string
	AdminKey          string
	dbQueries         database.Queries
	tokenDuration     int
	tokenSecret       string
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

type SleepSession struct {
	ID            uuid.UUID     `json:"id"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	SleepStart    time.Time     `json:"sleep_start"`
	SleepEnd      time.Time     `json:"sleep_end"`
	REMDuration   time.Duration `json:"rem_duration_mins"`
	LightDuration time.Duration `json:"light_duration_mins"`
	DeepDuration  time.Duration `json:"deep_duration_mins"`
	UserID        uuid.UUID     `json:"user_id"`
}

// --- End Struct Definitions

func main() {
	// Load Env Vars
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")

	// open connection to database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Failure to open connection to backend DB: %v", err)
		return
	}

	// Construct apiCfg for interfacing with env vars and DB
	var apiCfg apiCfg
	apiCfg.Port = os.Getenv("PORT")
	apiCfg.FilepathRoot = os.Getenv("FILEPATH_ROOT")
	apiCfg.UserCreationToken = os.Getenv("USER_CREATION_TOKEN")
	apiCfg.AdminKey = os.Getenv("ADMIN_KEY")
	apiCfg.dbQueries = *database.New(db)
	apiCfg.tokenDuration, _ = strconv.Atoi(os.Getenv("TOKEN_DURATION")) // Defaults to 3600 - one hour
	apiCfg.tokenSecret = os.Getenv("TOKEN_SECRET")

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
		// Check Access Token in Request
		token, err := auth.GetBearerToken(*req)
		if err != nil {
			log.Printf("Error parsing token: %v", err)
			http.Error(w, "Error parsing access token", http.StatusUnauthorized)
			return
		}
		if token != apiCfg.UserCreationToken {
			log.Printf("User creation request attempted w bad token")
			http.Error(w, "Access denied", http.StatusForbidden)
			return
		}

		// Set response header type
		w.Header().Set("Content-Type", "application/json")

		// Parse parameters from request
		type parameters struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		decoder := json.NewDecoder(req.Body)
		params := parameters{}
		err = decoder.Decode(&params)
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
			Email:        params.Email,
			PasswordHash: hashedPW,
		}

		// Create user in the DB
		userDB, err := apiCfg.dbQueries.CreateUser(req.Context(), dbParams)

		// Parse returned SQL data into user struct to return in response body
		user := User{
			ID:        userDB.ID,
			CreatedAt: userDB.CreatedAt,
			UpdatedAt: userDB.UpdatedAt,
			Email:     userDB.Email,
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

	// Login w/ User PW
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
		userDB, err := apiCfg.dbQueries.QueryUserEmail(req.Context(), params.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "User not found in DB", http.StatusNotFound)
				return
			}
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Check passwords match
		match, err := auth.CheckPasswordHash(params.Password, userDB.PasswordHash)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if !match {
			http.Error(w, "Incorrect Password", http.StatusUnauthorized)
			return
		}

		// Generate JWT - time.Duration defaults to nanoseconds (that was fun to learn)
		token, err := auth.MakeJWT(userDB.ID, apiCfg.tokenSecret, time.Duration(apiCfg.tokenDuration)*time.Second)
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
		passedToken, err := auth.GetBearerToken(*req)
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
			log.Printf("Database error: %v", err)
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
		token, err := auth.MakeJWT(refreshTokenDB.UserID, apiCfg.tokenSecret, time.Duration(apiCfg.tokenDuration)*time.Second)
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
		passedToken, err := auth.GetBearerToken(*req)
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

	// Reset the Users DB
	mux.HandleFunc("DELETE /api/users", func(w http.ResponseWriter, req *http.Request) {
		// Check Admin Key in headers
		token, err := auth.GetBearerToken(*req)
		if err != nil {
			log.Printf("Error parsing token: %v", err)
			http.Error(w, "Error parsing access token", http.StatusUnauthorized)
			return
		}
		if token != apiCfg.AdminKey {
			log.Printf("User creation request attempted w bad token")
			http.Error(w, "Access denied", http.StatusForbidden)
			return
		}

		err = apiCfg.dbQueries.RemoveUsers(req.Context())
		if err != nil {
			log.Printf("Error resetting users DB: %v", err)
			http.Error(w, "Error performing DB reset", http.StatusInternalServerError)
			return
		}

		log.Printf("Users DB has been reset")
		w.WriteHeader(http.StatusNoContent)
	})

	// - Sleep Functions -
	// Create Sleep Session
	mux.HandleFunc("POST /api/sleeps", func(w http.ResponseWriter, req *http.Request) {
		// Check Access Token in Request
		token, err := auth.GetBearerToken(*req)
		if err != nil {
			log.Printf("Error parsing token: %v", err)
			http.Error(w, "Error parsing access token", http.StatusUnauthorized)
			return
		}

		// Check Token is Valid - Returns user ID
		userDBID, err := auth.ValidateJWT(token, apiCfg.tokenSecret)
		if err != nil {
			if err == auth.ErrInvalidToken {
				log.Printf("Invalid token use attempted - need to use refresh")
				http.Error(w, "Error - Invalid token", http.StatusUnauthorized)
				return
			}
			log.Printf("Error processing token: %v", err)
			http.Error(w, "Error processing token", http.StatusUnauthorized)
			return
		}

		// TODO / Improvements: Find a way to handle generating a refresh token with
		// JWT expired

		// Parse out the parameters for the entry
		type parameters struct {
			SleepStart    *time.Time     `json:"sleep_start"`
			SleepEnd      *time.Time     `json:"sleep_end"`
			REMDuration   *time.Duration `json:"rem_duration_mins"`
			LightDuration *time.Duration `json:"light_duration_mins"`
			DeepDuration  *time.Duration `json:"deep_duration_mins"`
		}

		decoder := json.NewDecoder(req.Body)
		params := parameters{}
		err = decoder.Decode(&params)
		if err != nil {
			log.Printf("Error decoding parameters: %v", err)
			http.Error(w, "Error decoding parameters", http.StatusBadRequest)
			return
		}

		// Check that all values exists and are not empty
		if params.SleepStart == nil || params.SleepStart.IsZero() {
			log.Printf("Error - sleep start required in request body")
			http.Error(w, "Sleep start required", http.StatusBadRequest)
			return
		}
		if params.SleepEnd == nil || params.SleepEnd.IsZero() {
			log.Printf("Error - sleep end required in request body")
			http.Error(w, "Sleep end required", http.StatusBadRequest)
			return
		}
		if !params.SleepEnd.After(*params.SleepStart) {
			log.Printf("Error - sleep end must be after sleep start")
			http.Error(w, "Sleep end must be after sleep start", http.StatusBadRequest)
			return
		}
		if params.REMDuration == nil {
			log.Printf("Error - REM duration required in request body")
			http.Error(w, "REM duration required", http.StatusBadRequest)
			return
		}
		if params.LightDuration == nil {
			log.Printf("Error - light duration required in request body")
			http.Error(w, "Light duration required", http.StatusBadRequest)
			return
		}
		if params.DeepDuration == nil {
			log.Printf("Error - deep duration required in request body")
			http.Error(w, "Deep duration required", http.StatusBadRequest)
			return
		}
		if *params.REMDuration < 0 || *params.LightDuration < 0 || *params.DeepDuration < 0 {
			log.Printf("Error - sleep stage durations cannot be negative")
			http.Error(w, "Sleep stage durations cannot be negative", http.StatusBadRequest)
			return
		}

		// Parse request body parameters into DB struct
		databaseParams := database.CreateSleepSessionParams{
			SleepStart:        *params.SleepStart,
			SleepEnd:          *params.SleepEnd,
			RemDurationMins:   int32(*params.REMDuration),
			LightDurationMins: int32(*params.LightDuration),
			DeepDurationMins:  int32(*params.DeepDuration),
			UserID:            userDBID,
		}

		// Push data to the DB
		data, err := apiCfg.dbQueries.CreateSleepSession(req.Context(), databaseParams)
		if err != nil {
			log.Printf("Error writing sleep session to the DB: %v", err)
			http.Error(w, "Error writing sleep session to the DB", http.StatusInternalServerError)
			return
		}

		// Parse data to the return struct
		dat, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error marshaling sleep session to JSON: %v", err)
			http.Error(w, "Error marshaling sleep session to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(dat)

	})

	// Get sleep data by user ID
	mux.HandleFunc("GET /api/sleeps", func(w http.ResponseWriter, req *http.Request) {
		// Check Access Token in Request
		token, err := auth.GetBearerToken(*req)
		if err != nil {
			log.Printf("Error parsing token: %v", err)
			http.Error(w, "Error parsing access token", http.StatusUnauthorized)
			return
		}

		// Check Token is Valid - Returns user ID
		userDBID, err := auth.ValidateJWT(token, apiCfg.tokenSecret)
		if err != nil {
			if err == auth.ErrInvalidToken {
				log.Printf("Invalid token use attempted - need to use refresh")
				http.Error(w, "Error - Invalid token", http.StatusUnauthorized)
				return
			}
			log.Printf("Error processing token: %v", err)
			http.Error(w, "Error processing token", http.StatusUnauthorized)
			return
		}

		// Query DB by user ID
		DBData, err := apiCfg.dbQueries.QueryUserSleepSessions(req.Context(), userDBID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Printf("Recieved query for user data that does not exist")
				http.Error(w, "User does not have sessions recorded", http.StatusNotFound)
				return
			}
			log.Printf("Error getting user data from DB: %v", err)
			http.Error(w, "Error getting User Data", http.StatusInternalServerError)
			return
		}

		// Parse the data from the DB into a slice of sessions
		var sleepSessions []SleepSession

		for _, session := range DBData {
			sesh := SleepSession{
				ID:            session.ID,
				CreatedAt:     session.CreatedAt,
				UpdatedAt:     session.UpdatedAt,
				SleepStart:    session.SleepStart,
				SleepEnd:      session.SleepEnd,
				REMDuration:   time.Duration(session.RemDurationMins),
				LightDuration: time.Duration(session.LightDurationMins),
				DeepDuration:  time.Duration(session.DeepDurationMins),
				UserID:        session.UserID,
			}
			sleepSessions = append(sleepSessions, sesh)
		}

		// Marshal data and return
		data, err := json.Marshal(sleepSessions)
		if err != nil {
			log.Printf("Error marshaling sleep data: %v", err)
			http.Error(w, "Error marshaling sleep data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write(data)
	})

	// Reset the Sleep DB
	mux.HandleFunc("DELETE /api/sleeps", func(w http.ResponseWriter, req *http.Request) {
		// Check Admin Key in headers
		token, err := auth.GetBearerToken(*req)
		if err != nil {
			log.Printf("Error parsing token: %v", err)
			http.Error(w, "Error parsing access token", http.StatusUnauthorized)
			return
		}
		if token != apiCfg.AdminKey {
			log.Printf("User creation request attempted w bad token")
			http.Error(w, "Access denied", http.StatusForbidden)
			return
		}

		err = apiCfg.dbQueries.RemoveSleeps(req.Context())
		if err != nil {
			log.Printf("Error resetting sleep DB: %v", err)
			http.Error(w, "Error performing DB reset", http.StatusInternalServerError)
			return
		}

		log.Printf("Sleep DB has been reset")
		w.WriteHeader(http.StatusNoContent)
	})

	// --- End API Endpoint Definitions

	// Start Server
	server := &http.Server{
		Addr:    apiCfg.Port,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
