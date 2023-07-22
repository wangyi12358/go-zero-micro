package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-micro/api/admin/internal/logic/admin"
	"go-zero-micro/api/admin/internal/svc"
	"go-zero-micro/api/admin/internal/types"
)

func SendMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendMessageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewSendMessageLogic(r.Context(), svcCtx)
		err := l.SendMessage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
