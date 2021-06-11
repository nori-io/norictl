generate: ## generate protobuf
	protoc --proto_path=./api/protobuf -I=api/protobuf/  --go_out=plugins=grpc:. api/protobuf/*.proto
.PHONY: generate