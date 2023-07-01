.PHONY: protos

protos:
	protoc -I protos/ protos/currency.proto --go_out=protos/currency
	protoc -I protos/ protos/currency.proto --go-grpc_out=protos/currency