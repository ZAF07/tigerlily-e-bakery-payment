gen_wire:
	wire ./internal/injection

start:
	@echo "~~~~~~ Starting payment server ~~~~~~"
	@go run cmd/main.go