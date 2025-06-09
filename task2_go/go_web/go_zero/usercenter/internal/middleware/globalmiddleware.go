package middleware

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strconv"
)

type SetUidToCtxMiddleware struct {
}

func NewSetUidToCtxMiddleware() *SetUidToCtxMiddleware {
	return &SetUidToCtxMiddleware{}
}

func (m *SetUidToCtxMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 在中间件中将 header 中的 User-Agent 信息存到 context中
		logx.Info("SetUidToCtxMiddleware before")
		userID, err := strconv.ParseInt(r.Header.Get("X-User"), 10, 64)
		if err != nil {
			logx.Info("缺少id.")
			// 1. 直接写入错误响应
			w.WriteHeader(http.StatusOK) // 设置 HTTP 400 状态码
			fmt.Fprintf(w, `{"code":500, "msg": "缺少用户ID"}`)

			// 2. 提前返回，不调用 next
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", userID)

		next(w, r.WithContext(ctx))
		logx.Info("SetUidToCtxMiddleware after")
	}
}
