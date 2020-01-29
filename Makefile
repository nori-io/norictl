generate: ## generate protobuf
	@protoc --proto_path=api/protobuf/common_messages -I=api/protobuf/ --go_out=plugins=grpc:./internal/generated/protobuf/common_messages api/protobuf/common_messages/*.proto
	@protoc --proto_path=api/protobuf/config_messages -I=api/protobuf/ --go_out=plugins=grpc:./internal/generated/protobuf/config_messages api/protobuf/config_messages/*.proto
	@protoc --proto_path=api/protobuf/plugin_messages -I=api/protobuf/ --go_out=plugins=grpc:./internal/generated/protobuf/plugin_messages api/protobuf/plugin_messages/*.proto
	@protoc --proto_path=./api/protobuf/plugin_messages -I=api/protobuf/  --go_out=plugins=grpc:./internal/generated/protobuf/ api/protobuf/*.proto
.PHONY: generate 