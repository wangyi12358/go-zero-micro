package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-micro/api/admin/internal/logic"
	"go-zero-micro/api/admin/internal/svc"
	"go-zero-micro/api/admin/internal/types"
	"go-zero-micro/common/response"
	"net/http"
)

func sendMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendMessageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSendMessageLogic(r.Context(), svcCtx)
		err := l.SendMessage(&req)
		response.Response(w, nil, err) //②

	}
}
