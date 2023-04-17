gen_rpc_code:
	goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. --style=go_zero

start:
	go run user.go -f etc/user.yaml