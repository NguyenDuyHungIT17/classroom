package logic

import (
	"context"
	"strings"

	"backend/common"
	"backend/service/classroom/api/internal/svc"
	"backend/service/classroom/api/internal/types"
	"backend/service/classroom/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// GetUsers
func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersLogic) GetUsers(req *types.GetUsersReq) (resp *types.GetUsersRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Infof("GetUsers: %v", req)

	var users []*model.Users
	var usersOutput []types.User

	mapConditions := map[string]interface{}{
		"email":  strings.TrimSpace(req.Email),
		"limit":  req.Limit,
		"offset": req.Offset,
	}
	users, err = l.svcCtx.UsersModel.FindMultipleByConditionWithPagging(l.ctx, mapConditions)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetUsersRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	if len(users) == 0 {
		l.Logger.Error(err)
		return &types.GetUsersRes{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
			Data: types.GetUsersData{
				User: []types.User{},
			},
		}, nil
	}

	for _, user := range users {

		usersOutput = append(usersOutput, types.User{
			Id:               user.Id,
			UserName:         user.UserName,
			Password:         user.Password,
			Email:            user.Email,
			PhoneNumber:      user.PhoneNumber.String,
			Gender:           int(user.Gender),
			FullName:         user.FullName,
			Avatar:           user.Avatar.String,
			IsVerified:       user.IsVerified,
			VerificationCode: user.VerificationCode.String,
			Role:             int(user.Role),
			CreateTime:       user.CreateTime,
			UpdateTime:       user.UpdateTime,
		})
	}

	resp = &types.GetUsersRes{
		Code:    common.SUCCESS_CODE,
		Message: common.SUCCESS_MESS,
		Data: types.GetUsersData{
			User: usersOutput,
		},
	}

	l.Logger.Infof("GetUsers success: %v", resp)
	return resp, nil
}
