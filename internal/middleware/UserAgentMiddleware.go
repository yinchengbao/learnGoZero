package middleware

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
)

type UserAgentMiddleware struct {
}

func NewUserAgentMiddleware() *UserAgentMiddleware {
	return &UserAgentMiddleware{}
}

func (m *UserAgentMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		jwtToken := r.Header.Get("Authorization")
		if jwtToken == "" && strings.Index(r.URL.Path, "login") < 0 {
			httpx.ErrorCtx(r.Context(), w, errors.New("token 获取失败"))
		} else {
			next(w, r)
		}

	}
}
