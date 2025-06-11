package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"smart-api/internal/auth"
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
		http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
		return
	}

	claims, err := ExtractClaimsFromToken(token)
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка извлечения claims: %v", err), http.StatusUnauthorized)
		return
	}
	
	err = auth.RegisterUser(claims.Sub, claims.PreferredUsername, " ")
	if err != nil{
		http.Error(w, "ошибка при добавлении персоны", http.StatusUnauthorized)
		fmt.Println("jj")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token":    token,
		"username": claims.PreferredUsername,
		"userId":   claims.Sub,
	})
}
