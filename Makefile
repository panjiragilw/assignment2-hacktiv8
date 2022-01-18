server:
	nodemon --exec go run main.go --signal SIGTERM

.PHONY: sqlc server