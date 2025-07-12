package handler

import (
	"net/http"

	"backend/service/classroom/api/internal/logic"
	"backend/service/classroom/api/internal/svc"
	"backend/service/classroom/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// AddEnrollment
func AddEnrollmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddEnrollmentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddEnrollmentLogic(r.Context(), svcCtx)
		resp, err := l.AddEnrollment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
