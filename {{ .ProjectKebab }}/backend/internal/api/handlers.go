package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

// RegisterRoutes registers all API routes
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/health", healthHandler)
	mux.HandleFunc("/api/hello", helloHandler)
}

// healthHandler responds with a simple health check
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
	})
}

// helloHandler responds with a greeting
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := Response{
		Message: "Hello from {{ .Project }} API!",
		Time:    time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
