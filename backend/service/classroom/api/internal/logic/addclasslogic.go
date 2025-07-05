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
	"backend/service/classroom/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// AddClass
func NewAddClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassLogic {
	return &AddClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddClassLogic) AddClass(req *types.AddClassReq) (resp *types.AddClassRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Infof("AddClass: %v", req)

	currentTime := time.Now().UnixMilli()
	var classCode string
	var teacher *model.Teachers

	for i := 0; i < 5; i++ {
		code := utils.GenerateClassCode()
		exist, err := l.svcCtx.ClassesModel.ExistsByClassCode(l.ctx, code)
		if err != nil {
			l.Logger.Error(err)
			return &types.AddClassRes{
				Code:    common.DB_ERROR_CODE,
				Message: common.DB_ERROR_MESS,
			}, nil
		}
		if !exist {
			classCode = code
			break
		}
	}

	teacherId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil || teacherId == 0 {
		l.Logger.Error(common.INVALID_SESSION_USER_MESS)
		return &types.AddClassRes{
			Code:    common.INVALID_SESSION_USER_CODE,
			Message: common.INVALID_SESSION_USER_MESS,
		}, nil
	}

	newClass := &model.Classes{
		ClassCode:   classCode,
		ClassName:   req.ClassName,
		Description: sql.NullString{String: req.Description, Valid: true},
		TeacherId:   teacherId,
		CreateTime:  currentTime,
		UpdateTime:  currentTime,
	}

	classId, err := l.svcCtx.ClassesModel.InsertDb(l.ctx, newClass)
	if err != nil {
		l.Logger.Error(common.DB_ERROR_MESS)
		return &types.AddClassRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	newClass.Id = classId
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	//add teacher after create class
	bio := newClass.ClassName + ": " + newClass.Description.String
	teacher = &model.Teachers{
		UserId:     userId,
		Bio:        sql.NullString{String: bio, Valid: true},
		CreateTime: currentTime,
		UpdateTime: currentTime,
		ClassId:    sql.NullInt64{Int64: classId, Valid: true},
	}

	err = l.svcCtx.TeachersModel.InsertDb(l.ctx, teacher)
	if err != nil {
		l.Logger.Error(common.ADD_TEACHER_ERROR_MESS)
		return &types.AddClassRes{
			Code:    common.ADD_TEACHER_ERROR_CODE,
			Message: common.ADD_TEACHER_ERROR_MESS,
		}, nil
	}

	resp = &types.AddClassRes{
		Code:    common.SUCCESS_CODE,
		Message: common.SUCCESS_MESS,
		Data: types.AddClassData{
			Class: types.Class{
				Id:          newClass.Id,
				ClassCode:   newClass.ClassCode,
				ClassName:   newClass.ClassName,
				Description: newClass.Description.String,
				TeacherId:   newClass.TeacherId,
				CreateTime:  newClass.CreateTime,
				UpdateTime:  newClass.UpdateTime,
			},
			Teacher: types.Teacher{
				UserId:     teacher.UserId,
				Bio:        teacher.Bio.String,
				CreateTime: teacher.CreateTime,
				UpdateTime: teacher.UpdateTime,
				ClassId:    classId,
			},
		},
	}
	return
}
