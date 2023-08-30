package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"login/internal/logic"
	"login/internal/types"
	"net/http"

	"login/internal/svc"
)

func MenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		l := logic.NewMenuLogic(r.Context(), svcCtx)
		resp, err := l.MenuLogic(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
