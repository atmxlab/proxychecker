HTTP_CMD_PATH=./cmd/http
CONSOLE_CMD_PATH=./cmd/console

run\:console:
	go run $(CONSOLE_CMD_PATH)

run\:http:
	go run $(HTTP_CMD_PATH)

mockgen\:install:
	go install github.com/golang/mock/mockgen@latest

generate:
	go generate ./...

