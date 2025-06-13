package main

import (
	"net/http"
	"smart-api/api"
	"smart-api/internal/middleware"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/login", api.LoginHandler)
	mux.HandleFunc("/api/charges", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			api.ChargesHandler(w, r)
		} else if r.Method == http.MethodPost {
			api.CreateChargeHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/charges/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			api.ChargeHandlerByID(w, r)
		}
		if r.Method == http.MethodPut {
			api.ChargesHandlerPut(w, r)
		}
	})
	mux.HandleFunc("/api/accounts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			api.AccountsHadlerGet(w, r)
		} else if r.Method == http.MethodPost {
			api.AccountsHadlerPost(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/accounts/", api.AccountsHandlerByUserID)
	mux.HandleFunc("/api/register", api.RegisterHandler)

	handler := middleware.CorsMiddleware(mux)
	http.ListenAndServe(":8080", handler)
}
