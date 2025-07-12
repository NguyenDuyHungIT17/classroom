package logic

import (
	"context"
	"encoding/json"
	"time"

	"backend/common"
	"backend/service/classroom/api/internal/svc"
	"backend/service/classroom/api/internal/types"
	"backend/service/classroom/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddEnrollmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// AddEnrollment
func NewAddEnrollmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddEnrollmentLogic {
	return &AddEnrollmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddEnrollmentLogic) AddEnrollment(req *types.AddEnrollmentReq) (resp *types.AddEnrollmentRes, err error) {
	// todo: add your logic here and delete this line

	var class *model.Classes
	currentTime := time.Now().UnixMilli()

	class, err = l.svcCtx.ClassesModel.FindOneByClassCode(l.ctx, req.ClassCode)
	if class == nil || err != nil {
		l.Logger.Error(common.CLASS_IS_NOT_EXIST_MESS)
		return &types.AddEnrollmentRes{
			Code:    common.CLASS_IS_NOT_EXIST_CODE,
			Message: common.CLASS_IS_NOT_EXIST_MESS,
			Data:    types.AddEnrollmentData{},
		}, nil
	}

	userIdAtToken, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil || userIdAtToken == 0 {
		l.Logger.Error(common.INVALID_SESSION_USER_MESS)
		return &types.AddEnrollmentRes{
			Code:    common.INVALID_SESSION_USER_CODE,
			Message: common.INVALID_SESSION_USER_MESS,
		}, nil
	}

	exist, err := l.svcCtx.EnrollmentsModel.ExistsByClassIdAndStudentId(l.ctx, class.Id, userIdAtToken)
	if err != nil {
		l.Logger.Error(common.DB_ERROR_MESS)
		return &types.AddEnrollmentRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	if !exist {

		enrollment := &model.Enrollments{
			ClassId:   class.Id,
			StudentId: userIdAtToken,
			JoinTime:  currentTime,
		}

		err = l.svcCtx.EnrollmentsModel.InsertDb(l.ctx, enrollment)
		if err != nil {
			l.Logger.Error(common.DB_ERROR_MESS)
			return &types.AddEnrollmentRes{
				Code:    common.DB_ERROR_CODE,
				Message: common.DB_ERROR_MESS,
			}, nil
		}

		resp = &types.AddEnrollmentRes{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
			Data: types.AddEnrollmentData{
				Enrollment: types.Enrollment{
					Id:        enrollment.Id,
					StudentId: enrollment.StudentId,
					ClassId:   enrollment.ClassId,
					JoinTime:  enrollment.JoinTime,
				},
			},
		}

		l.Logger.Infof("addEnrollments Success: %v", resp)
	}
	return
}
