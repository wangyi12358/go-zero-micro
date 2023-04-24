package handler

import (
	"go-zero-micro/api/admin/internal/logic"
	"go-zero-micro/api/admin/internal/svc"
	"go-zero-micro/common/response"
	"net/http"
)

func profileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewProfileLogic(r.Context(), svcCtx)
		resp, err := l.Profile()
		response.Response(w, resp, err) //②

	}
}
