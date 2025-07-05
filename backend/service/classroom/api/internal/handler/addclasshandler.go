package handler

import (
	"net/http"

	"backend/service/classroom/api/internal/logic"
	"backend/service/classroom/api/internal/svc"
	"backend/service/classroom/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// AddClass
func AddClassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddClassReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddClassLogic(r.Context(), svcCtx)
		resp, err := l.AddClass(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
