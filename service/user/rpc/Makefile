gen_rpc_code:
	goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

start:
	go run user.go -f etc/user.yaml