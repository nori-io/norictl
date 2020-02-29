generate: ## generate protobuf
	@protoc --proto_path=api/protobuf/common -I=api/protobuf/ --go_out=plugins=grpc:./internal/generated/protobuf/common api/protobuf/common/*.proto
	@protoc --proto_path=api/protobuf/config -I=api/protobuf/ --go_out=plugins=grpc:./internal/generated/protobuf/config api/protobuf/config/*.proto
	@protoc --proto_path=api/protobuf/plugin -I=api/protobuf/ --go_out=plugins=grpc:./internal/generated/protobuf/plugin api/protobuf/plugin/*.proto
	@protoc --proto_path=./api/protobuf/plugin -I=api/protobuf/  --go_out=plugins=grpc:./internal/generated/protobuf/ api/protobuf/*.proto
.PHONY: generate 