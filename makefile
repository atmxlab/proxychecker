CMD_PATH=./cmd
run:
	go run $(CMD_PATH)

mockgen\:install:
	go install github.com/golang/mock/mockgen@latest

generate:
	go generate ./...