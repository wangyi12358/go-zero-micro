package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-micro/api/admin/internal/logic/user"
	"go-zero-micro/api/admin/internal/svc"
	"go-zero-micro/api/admin/internal/types"
)

func UserPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserPageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserPageLogic(r.Context(), svcCtx)
		resp, err := l.UserPage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
