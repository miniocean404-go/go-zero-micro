package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero/apps/app/api/internal/logic/user"
	"go_zero/apps/app/api/internal/svc"
	"go_zero/apps/app/api/internal/types"
)

func UserCollectionDelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCollectionDelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserCollectionDelLogic(r.Context(), svcCtx)
		resp, err := l.UserCollectionDel(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
