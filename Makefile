generate: ## generate protobuf
	@protoc --proto_path=nori-grpc/plugin --go_out=plugins=grpc:./internal/generated/protobuf/plugin nori-grpc/plugin/*.proto
	@protoc -I. --proto_path=nori-grpc/plugin --go_out=plugins=grpc:./internal/generated/protobuf/ nori-grpc/*.proto
.PHONY: generate 
