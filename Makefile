export ETCD_ADDR="127.0.0.1:2379"

gen_db_model:
	go run ./script/gen_db_model.go

start_user_rpc:
	go run ./service/user/rpc/user.go -f ./service/user/rpc/etc/user.yaml

start_admin_api:
	go run ./api/admin/admin.go -f ./api/admin/etc/admin.yaml

push_etcd_config:
	go run ./script/push_etcd_config.go 11324
