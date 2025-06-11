package api

import (
	"encoding/json"
	"net/http"
	"smart-api/internal/auth"
	"smart-api/internal/httpx"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
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

	user, err := auth.LoginUser(req.Username, req.Password)
	if err != nil {
		httpx.NewJSONError(w, http.StatusUnauthorized, "Invalid credentials", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token":    "fake-token-123",
		"username": user.Username,
		"userId":   user.ID,
	})
}
