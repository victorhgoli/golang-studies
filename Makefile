# Makefile

proto:
	@echo "Gerando arquivos protobuf..."
	@protoc --go_out=. --go-grpc_out=.  grpc/service.proto



mockgen -source=internal/repository/contract.go  -destination=tests/mocks/repository/repository.go


mockgen -source=infra/http_client/http_client.go -destination=tests/mocks/http_client/http_client.go


mockgen -source=integration/contract.go  -destination=tests/mocks/integration/integration.go