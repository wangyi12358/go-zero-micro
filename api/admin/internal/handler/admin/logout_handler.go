package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-micro/api/admin/internal/logic/admin"
	"go-zero-micro/api/admin/internal/svc"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewLogoutLogic(r.Context(), svcCtx)
		err := l.Logout()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
