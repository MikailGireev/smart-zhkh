.PHONY: all smart-api smart-zhkh auth-api run kill

all: run

run: smart-api smart-zhkh auth-api admin-charge admin-task

smart-api:
	cd smart-api && go run cmd/main.go &

smart-zhkh:
	cd smart-zhkh && npm install && npm run dev &

auth-api:
	cd auth-api && go run cmd/main.go &

admin-charge:
	cd admin-api/charge_service && go run cmd/main.go &

admin-task:
	cd admin-api/task_service && go run cmd/main.go &

kill:
	-pkill -f "go run" || true
	-pkill -f "npm" || true
