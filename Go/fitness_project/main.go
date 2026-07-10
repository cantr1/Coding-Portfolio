package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)


type apiCfg struct {
	Port string
	FilepathRoot string
}

func main() {
	// Load Env Vars
	godotenv.Load()
	var apiCfg apiCfg
	apiCfg.Port = os.Getenv("PORT")
	apiCfg.FilepathRoot = os.Getenv("FILEPATH_ROOT")

	// Create a multiplexer
	var mux = http.NewServeMux()

	// --- Begin API Endpoint Definitions

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
			http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	// --- End API Endpoint Definitions
	
	// Start Server
	server := &http.Server{
		Addr: apiCfg.Port,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}