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

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// UpdateUser
func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.UpdateUserRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Infof("UpdateUser: %v", req)

	var userId int64
	var user *model.Users
	currentTime := time.Now().UnixMilli()

	userId, err = l.ctx.Value("userId").(json.Number).Int64()
	if userId == 0 || err != nil {
		l.Logger.Error(common.INVALID_SESSION_USER_MESS)
		return &types.UpdateUserRes{
			Code:    common.INVALID_SESSION_USER_CODE,
			Message: common.INVALID_SESSION_USER_MESS,
		}, nil
	}

	user, err = l.svcCtx.UsersModel.FindOne(l.ctx, userId)
	if user == nil || err != nil {
		l.Logger.Error(common.DB_ERROR_MESS)
		return &types.UpdateUserRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	if user.Id != userId {
		l.Logger.Error(common.INVALID_SESSION_USER_MESS)
		return &types.UpdateUserRes{
			Code:    common.INVALID_SESSION_USER_CODE,
			Message: common.INVALID_SESSION_USER_MESS,
		}, nil
	}

	user.PhoneNumber = sql.NullString{String: req.PhoneNumber, Valid: true}
	user.Gender = int64(req.Gender)
	user.FullName = req.FullName
	user.UpdateTime = currentTime
	if err = l.svcCtx.UsersModel.Update(l.ctx, user); err != nil {
		l.Logger.Error(err)
		return &types.UpdateUserRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	userOutput := types.User{
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
	}

	resp = &types.UpdateUserRes{
		Code:    common.SUCCESS_CODE,
		Message: common.SUCCESS_MESS,
		Data: types.UpdateUserData{
			User: userOutput,
		},
	}

	l.Logger.Infof("Update User success: %v", resp)
	return resp, nil
}
