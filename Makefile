
gen_db_model:
	go run ./script/gen_db_model.go

gen_user_rpc:
	goctl rpc protoc ./service/user/rpc/user.proto --go_out=./service/user/rpc/types --go-grpc_out=./service/user/rpc/types --zrpc_out=./service/user/rpc/

start_user_rpc:
	go run ./service/user/rpc/user.go -f ./service/user/rpc/etc/user.yaml

start_gateway_api:
	go run ./service/gateway/api/gateway.go -f ./service/gateway/api/etc/gateway.yaml
