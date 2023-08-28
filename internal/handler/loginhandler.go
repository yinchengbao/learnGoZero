package handler

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"login/internal/logic"
	"login/internal/svc"
	"login/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if resp == nil {
			resp = &types.Response{}
		}

		result, err := svcCtx.UserModel.FindALL(r.Context())
		if err != nil {
			fmt.Println("报错了")
		}
		for i, user := range result {
			fmt.Println("查询数据库中得数据")
			fmt.Println(i, user.Name)
		}

		resp.Message = "返回消息"
		resp.Users = result
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
