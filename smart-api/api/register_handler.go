package api

import (
	"encoding/json"
	"net/http"
	"smart-api/internal/auth"
	"smart-api/internal/httpx"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return 
	}

	var req auth.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.NewJSONError(w, http.StatusBadRequest, "Invalid request", "Invalid request")
		return
	}

	if req.Username == "" || req.Password == "" {
		httpx.NewJSONError(w, http.StatusBadRequest, "Invalid request", "Username and password are required")
		return
	}

	err := auth.RegisterUser(req.Username, req.Password)
	if err != nil {
		httpx.NewJSONError(w, http.StatusInternalServerError, "User regsistered successfully", err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User regsistered successfully"))
}