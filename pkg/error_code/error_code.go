package error_code

import "go-zero-micro/pkg/response"

var (
	UserNotExist = response.Response{Code: 1000, Msg: "账号或密码错误!"}
)
