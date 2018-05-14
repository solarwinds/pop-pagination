GO_RPCDIR := go-rpc

pbuf_go:
	protoc -I. --go_out=plugins=grpc:$(GO_RPCDIR) *.proto