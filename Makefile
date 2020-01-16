generate: ## generate protobuf
	@protoc --proto_path=api/protobuf/plugin --go_out=plugins=grpc:./internal/generated/protobuf/plugin api/protobuf/plugin/*.proto
	@protoc --proto_path=./api/protobuf/plugin -I=api/protobuf/ --go_out=plugins=grpc:./internal/generated/protobuf/ api/protobuf/*.proto
.PHONY: generate 