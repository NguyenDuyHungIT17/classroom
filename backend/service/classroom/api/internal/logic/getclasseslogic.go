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
	l.Logger.Infof("GetUsers: %v", req)

	var classes []*model.Classes
	var classesOutput []types.Class

	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if userId == 0 || err != nil {
		l.Logger.Error(common.INVALID_SESSION_USER_MESS)
		return &types.GetClassesRes{
			Code:    common.INVALID_SESSION_USER_CODE,
			Message: common.INVALID_SESSION_USER_MESS,
		}, nil
	}

	mapConditions := map[string]interface{}{
		"limit":  req.Limit,
		"offset": req.Offset,
	}
	classes, err = l.svcCtx.ClassesModel.FindMultipleByConditionWithPagging(l.ctx, mapConditions, userId)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetClassesRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	if len(classes) == 0 {
		l.Logger.Error(err)
		return &types.GetClassesRes{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
			Data: types.GetClassesData{
				Class: []types.Class{},
			},
		}, nil
	}

	for _, class := range classes {
		classesOutput = append(classesOutput, types.Class{
			ClassCode:   class.ClassCode,
			ClassName:   class.ClassName,
			Description: class.Description.String,
			CreateTime:  class.CreateTime,
		})
	}

	resp = &types.GetClassesRes{
		Code:    common.SUCCESS_CODE,
		Message: common.SUCCESS_MESS,
		Data: types.GetClassesData{
			Class: classesOutput,
		},
	}

	l.Logger.Infof("GetClasses success: %v", resp)
	return resp, nil
}
