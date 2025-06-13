package api

import (
	"encoding/json"
	"net/http"
	"smart-api/internal/auth"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	token, err := auth.LoginUserKeycloak(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials: "+err.Error(), http.StatusUnauthorized)
		return
	}

	claims, err := auth.ExtractClaimsFromToken(token.AccessToken)
	if err != nil {
		http.Error(w, "Failed to extract token claims: "+err.Error(), http.StatusUnauthorized)
		return
	}
<<<<<<< HEAD:smart-zhkh/smart-api/api/login_handler.go
=======
	
	err = auth.RegisterUser(claims.Sub, claims.PreferredUsername, " ")
	if err != nil{
		http.Error(w, "ошибка при добавлении персоны", http.StatusUnauthorized)
	}
>>>>>>> 46748f6b97a3be5c70c261d8560b8ca2c3646174:smart-api/api/login_handler.go

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token":    token.AccessToken,
		"userId":   claims.Sub,
		"username": claims.PreferredUsername,
	})
}
