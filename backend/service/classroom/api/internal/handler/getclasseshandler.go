package handler

import (
	"net/http"

	"backend/service/classroom/api/internal/logic"
	"backend/service/classroom/api/internal/svc"
	"backend/service/classroom/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// GetClasses
func GetClassesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetClassesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetClassesLogic(r.Context(), svcCtx)
		resp, err := l.GetClasses(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
