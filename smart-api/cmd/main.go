package main

import (
	"fmt"
	"net/http"
	"smart-api/api"
	"smart-api/internal/middleware"
)

func main() {
	mux := http.NewServeMux()
	fmt.Print(12)
	mux.HandleFunc("/api/login", api.LoginHandler)
	mux.HandleFunc("/api/charges", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			api.ChargesHandler(w, r)
		} else if r.Method == http.MethodPost {
			api.CreateChargeHandler(w, r)
		}else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}) 

	mux.HandleFunc("/api/charges/", api.ChargeHandlerByID)
	mux.HandleFunc("/api/accounts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			api.AccountsHadlerGet(w, r)
		} else if r.Method == http.MethodPost {
			api.AccountsHadlerPost(w, r)
		}else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/accounts/", api.AccountsHandlerByUserID)

	handler := middleware.CorsMiddleware(mux)
	http.ListenAndServe(":8080", handler)
}