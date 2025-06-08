package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero/usercenter/internal/logic/user"
	"go_zero/usercenter/internal/svc"
	"go_zero/usercenter/internal/types"
)

// 新增用户
func UserCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserCreateLogic(r.Context(), svcCtx)
		resp, err := l.UserCreate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
