package logic

import (
	"context"

	"backend/service/classroom/api/internal/svc"
	"backend/service/classroom/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// GetClasses
func NewGetClassesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassesLogic {
	return &GetClassesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClassesLogic) GetClasses(req *types.GetClassesReq) (resp *types.GetClassesRes, err error) {
	// todo: add your logic here and delete this line

	return
}
