rpc:
	goctl rpc protoc proto/api.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=go_zero
gen:
	protoc --go_out=. --go-grpc_out=. proto/api.proto