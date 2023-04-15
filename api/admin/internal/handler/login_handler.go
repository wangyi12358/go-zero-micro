package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-micro/api/admin/internal/logic"
	"go-zero-micro/api/admin/internal/svc"
	"go-zero-micro/api/admin/internal/types"
	"go-zero-micro/common/response"
	"net/http"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err) //②

	}
}
