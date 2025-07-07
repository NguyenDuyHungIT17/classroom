package logic

import (
	"context"

	"backend/common"
	"backend/service/classroom/api/internal/svc"
	"backend/service/classroom/api/internal/types"
	"backend/service/classroom/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// GetClass
func NewGetClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassLogic {
	return &GetClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClassLogic) GetClass(req *types.GetClassReq) (resp *types.GetClassRes, err error) {
	// todo: add your logic here and delete this line
	var class *model.Classes

	class, err = l.svcCtx.ClassesModel.FindOne(l.ctx, req.ClassId)
	if class == nil || err != nil {
		l.Logger.Error(common.DB_ERROR_MESS)
		return &types.GetClassRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	classOutPut := types.Class{
		ClassCode:   class.ClassCode,
		ClassName:   class.ClassName,
		Description: class.Description.String,
		CreateTime:  class.CreateTime,
	}

	resp = &types.GetClassRes{
		Code:    common.SUCCESS_CODE,
		Message: common.SUCCESS_MESS,
		Data: types.GetClassData{
			Class: classOutPut,
		},
	}
	l.Logger.Infof("Get Class Succes: %v", resp)
	return
}
