install_sqlc:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

run:
	go run cmd/main.go
