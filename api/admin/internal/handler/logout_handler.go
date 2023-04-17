package handler

import (
	_ "github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-micro/api/admin/internal/logic"
	"go-zero-micro/api/admin/internal/svc"
	"go-zero-micro/common/response"
	"net/http"
)

func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		err := l.Logout()
		response.Response(w, nil, err) //②

	}
}
