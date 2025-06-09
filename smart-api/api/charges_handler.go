package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"smart-api/internal/auth"
	"smart-api/internal/httpx"
)

func ChargesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodGet {
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	charges, err := auth.LoadCharges()
	if err != nil {
		fmt.Println("Failed to load charges:", err)
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to load charges", err.Error())
		return
	}

	json.NewEncoder(w).Encode(charges)
}

func CreateChargeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodPost {
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Method not allowed")
		return
	}

	var newCharge auth.Charge
	if err := json.NewDecoder(r.Body).Decode(&newCharge); err != nil {
		httpx.NewJSONError(w, http.StatusBadRequest, "Invalid request", "Invalid request")
		return
	}

	charges, err := auth.LoadCharges()
	if err != nil {
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to load charges", err.Error())
		return
	}

	newCharge.ID = len(charges) + 1
	charges = append(charges, newCharge)

	err = auth.SaveCharges(charges)
	if err != nil {
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to save charges", err.Error())
		return
	}

	json.NewEncoder(w).Encode(newCharge)
}