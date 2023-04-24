package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go-zero-micro/common/errorx"
	"go-zero-micro/service/user/rpc/types/user"
	"time"

	"go-zero-micro/api/admin/internal/svc"
	"go-zero-micro/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	userRes, err := l.svcCtx.UserRpc.UserLogin(l.ctx, &user.CheckUserRequest{
		Username: req.Username,
		Password: req.Passwrod,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("账号密码错误!")
	}
	token, err := l.getJwtToken(
		l.svcCtx.Config.Auth.AccessSecret,
		time.Now().Unix(),
		l.svcCtx.Config.Auth.AccessExpire,
		userRes.Id,
	)
	if err != nil {
		fmt.Printf("Error obtaining jwt token: %s", err.Error())
		return nil, err
	}
	return &types.LoginReply{
		Token:  token,
		Expire: l.svcCtx.Config.Auth.AccessExpire,
	}, nil
}

func GetUserId(ctx context.Context) int64 {
	jsonUserId := ctx.Value("userId").(json.Number)
	userId, _ := jsonUserId.Int64()
	return userId
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
