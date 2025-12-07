API_CMD_PATH=./cmd/api
CONSOLE_CMD_PATH=./cmd/console

run\:console:
	go run $(CONSOLE_CMD_PATH)

run\:api:
	go run $(API_CMD_PATH)

mockgen\:install:
	go install github.com/golang/mock/mockgen@latest

generate:
	go generate ./...

.PHONY: proto clean install-plugins

PROTO_DIR=api
GEN_DIR=gen/proto

proto\:install:
	@echo "Installing protoc plugins..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@echo "Plugins installed successfully"

proto\:gen:
	@# Проверяем наличие плагинов
	@command -v protoc-gen-go >/dev/null 2>&1 || { echo "protoc-gen-go not found, installing..."; go install google.golang.org/protobuf/cmd/protoc-gen-go@latest; }
	@command -v protoc-gen-go-grpc >/dev/null 2>&1 || { echo "protoc-gen-go-grpc not found, installing..."; go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest; }
	@command -v protoc-gen-grpc-gateway >/dev/null 2>&1 || { echo "protoc-gen-grpc-gateway not found, installing..."; go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest; }
	@command -v protoc-gen-openapiv2 >/dev/null 2>&1 || { echo "protoc-gen-openapiv2 not found, installing..."; go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest; }

	mkdir -p $(GEN_DIR)
	buf generate --path $(PROTO_DIR)

proto\:clean:
	rm -rf gen/proto