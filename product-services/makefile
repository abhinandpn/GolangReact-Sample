.PHONY: wire swag

wire: ./pkg/di/wire_gen.go ## Generate wire_gen.go
    cd pkg/di && wire

swag: ## Generate swagger docs
	swag init -g pkg/api/server.go -o ./cmd/api/docs
