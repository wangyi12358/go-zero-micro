// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `json:"username"`
	Passwrod string `json:"password"`
}

type LoginReply struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

type UserPageReq struct {
	Current  int16 `json:"current"`
	PageSize int16 `json:"pageSize"`
}

type CreateUserReq struct {
	Nickname string `json:"nickname"`
	Gender   int32  `json:"gender"`
	Username string `json:"username"`
	Passwrod string `json:"password"`
}

type UserPageReply struct {
	List  []UserReply `json:"list"`
	Total int32       `json:"total"`
}

type UserReply struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Gender   int32  `json:"gender"`
}
