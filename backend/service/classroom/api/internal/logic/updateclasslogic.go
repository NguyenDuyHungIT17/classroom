package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"backend/common"
	"backend/service/classroom/api/internal/svc"
	"backend/service/classroom/api/internal/types"
	"backend/service/classroom/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// UpdateClass
func NewUpdateClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClassLogic {
	return &UpdateClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateClassLogic) UpdateClass(req *types.UpdateClassReq) (resp *types.UpdateClassRes, err error) {
	// todo: add your logic here and delete this line
	var class *model.Classes
	var teacher *model.Teachers
	currentTime := time.Now().UnixMilli()
	class, err = l.svcCtx.ClassesModel.FindOne(l.ctx, req.ClassId)
	if class == nil || err != nil {
		l.Logger.Error(common.DB_ERROR_MESS)
		return &types.UpdateClassRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if userId == 0 || err != nil {
		l.Logger.Error(common.INVALID_SESSION_USER_MESS)
		return &types.UpdateClassRes{
			Code:    common.INVALID_SESSION_USER_CODE,
			Message: common.INVALID_SESSION_USER_MESS,
		}, nil
	}

	//check: teacher in class  is correct
	if class.TeacherId != userId {
		l.Logger.Error(common.INVALID_REQUEST_MESS)
		return &types.UpdateClassRes{
			Code:    common.INVALID_REQUEST_CODE,
			Message: common.INVALID_REQUEST_MESS,
		}, nil
	}

	teacher, err = l.svcCtx.TeachersModel.FindOneByUserId(l.ctx, userId)
	if teacher == nil || err != nil {
		l.Logger.Error(common.INVALID_REQUEST_MESS)
		return &types.UpdateClassRes{
			Code:    common.INVALID_REQUEST_CODE,
			Message: common.INVALID_REQUEST_MESS,
		}, nil
	}

	//update info class
	class.ClassName = req.ClassName
	class.Description = sql.NullString{String: req.Description, Valid: true}
	class.UpdateTime = currentTime

	//update info teacher
	bio := class.ClassName + ": " + class.Description.String
	teacher.Bio = sql.NullString{String: bio, Valid: true}
	teacher.UpdateTime = currentTime

	err = l.svcCtx.ClassesModel.Update(l.ctx, class)
	if err != nil {
		l.Logger.Error(common.DB_ERROR_MESS)
		return &types.UpdateClassRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	err = l.svcCtx.TeachersModel.Update(l.ctx, teacher)
	if err != nil {
		l.Logger.Error(common.DB_ERROR_MESS)
		return &types.UpdateClassRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	resp = &types.UpdateClassRes{
		Code:    common.SUCCESS_CODE,
		Message: common.SUCCESS_MESS,
		Data: types.UpdateClassData{
			Class: types.Class{
				Id:          class.Id,
				ClassCode:   class.ClassCode,
				ClassName:   class.ClassName,
				Description: class.Description.String,
				TeacherId:   class.TeacherId,
				CreateTime:  class.CreateTime,
				UpdateTime:  class.UpdateTime,
			},
		},
	}
	l.Logger.Infof("update class success: %v", resp)
	return
}
