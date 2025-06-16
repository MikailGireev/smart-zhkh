package api

import (
	"admin-api/task_service/internal/auth"
	"admin-api/task_service/internal/httpx"
	"admin-api/task_service/internal/logger"
	"encoding/json"
	"net/http"
)

var console = &logger.Console{}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		console.Errorf("Method %s not allowed", r.Method)
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only GET allowed")
		return
	}

	tasks, err := auth.LoadTasks()
	if err != nil {
		console.Errorf("Failed to load tasks: %v", err)
		httpx.NewJSONError(w, http.StatusInternalServerError, "Failed to load tasks", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}