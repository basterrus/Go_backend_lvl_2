proto:
	protoc ./pb/product.proto ./pb/pricelist.proto --go_out=./internal/proto --go-grpc_out=./internal/proto --proto_path=./pb

.PHONY: proto