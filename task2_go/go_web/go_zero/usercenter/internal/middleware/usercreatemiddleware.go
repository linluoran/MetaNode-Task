package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type UserCreateMiddleware struct {
}

func NewUserCreateMiddleware() *UserCreateMiddleware {
	return &UserCreateMiddleware{}
}

func (m *UserCreateMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 在中间件中将 header 中的 User-Agent 信息存到 context中
		logx.Info("UserCreateMiddleware before")
		val := r.Header.Get("User-Agent")
		reqCtx := r.Context()
		ctx := context.WithValue(reqCtx, "User-Agent", val)
		newReq := r.WithContext(ctx)
		next(w, newReq)
		logx.Info("UserCreateMiddleware after")
	}
}
