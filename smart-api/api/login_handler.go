package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

	token, err := GetToken(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Не правильный Логин или Пароль", http.StatusUnauthorized)
		return
	}

	username, err := ExtractUsernameFromToken(token)
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка извлечения username: %v", err), http.StatusUnauthorized)
		return
	}

	response := LoginResponse{
		Token:    token,
		Username: username,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}