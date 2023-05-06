
gen_db_model:
	go run ./script/gen_db_model.go

start_user_rpc:
	go run ./service/user/rpc/user.go -f ./service/user/rpc/etc/user.yaml

start_admin_api:
	go run ./api/admin/admin.go -f ./api/admin/etc/admin.yaml
