package api

import (
	"admin-api/task_service/internal/httpx"
	"admin-api/task_service/internal/logger"
	"net/http"
)

var console = &logger.Console{}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		console.Errorf("Method %s not allowed", r.Method)
		httpx.NewJSONError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only GET allowed")
		return
	}
}