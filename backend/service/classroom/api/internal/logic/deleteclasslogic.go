package logic

import (
	"context"
	"encoding/json"

	"backend/common"
	"backend/service/classroom/api/internal/svc"
	"backend/service/classroom/api/internal/types"
	"backend/service/classroom/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// DeleteClass
func NewDeleteClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteClassLogic {
	return &DeleteClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteClassLogic) DeleteClass(req *types.DeleteClassReq) (resp *types.DeleteClassRes, err error) {
	// todo: add your logic here and delete this line
	var class *model.Classes

	userIdAtToken, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil || userIdAtToken == 0 {
		l.Logger.Error(common.INVALID_SESSION_USER_MESS)
		return &types.DeleteClassRes{
			Code:    common.INVALID_SESSION_USER_CODE,
			Message: common.INVALID_SESSION_USER_MESS,
		}, nil
	}

	class, err = l.svcCtx.ClassesModel.FindOne(l.ctx, req.ClassId)
	if class == nil || err != nil {
		l.Logger.Error(common.DB_ERROR_MESS)
		return &types.DeleteClassRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	if userIdAtToken != class.TeacherId {
		l.Logger.Error(common.INVALID_REQUEST_MESS)
		return &types.DeleteClassRes{
			Code:    common.INVALID_REQUEST_CODE,
			Message: common.INVALID_REQUEST_MESS,
		}, nil
	}

	err = l.svcCtx.TeachersModel.DeleteByClassId(l.ctx, class.Id)
	if err != nil {
		l.Logger.Error(common.DB_ERROR_MESS)
		return &types.DeleteClassRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	err = l.svcCtx.EnrollmentsModel.DeleteByClassId(l.ctx, class.Id)
	if err != nil {
		l.Logger.Error(common.DB_ERROR_MESS)
		return &types.DeleteClassRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	err = l.svcCtx.ClassesModel.Delete(l.ctx, class.Id)
	if err != nil {
		l.Logger.Error(common.DB_ERROR_MESS)
		return &types.DeleteClassRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	resp = &types.DeleteClassRes{
		Code:    common.SUCCESS_CODE,
		Message: common.SUCCESS_MESS,
	}

	l.Logger.Infof("Delete Class success: %v", resp)
	return
}
