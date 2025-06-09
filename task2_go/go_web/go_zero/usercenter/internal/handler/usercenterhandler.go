package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"usercenter/internal/logic"
	"usercenter/internal/svc"
	"usercenter/internal/types"
)

func UsercenterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUsercenterLogic(r.Context(), svcCtx)
		resp, err := l.Usercenter(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
