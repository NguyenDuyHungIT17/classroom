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

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// GetUser
func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Infof("Get User: %v", req)

	var userId int64
	var user *model.Users

	userId, err = l.ctx.Value("userId").(json.Number).Int64()
	if err != nil || userId == 0 {
		l.Logger.Error(common.INVALID_SESSION_USER_CODE)
		return &types.GetUserRes{
			Code:    common.INVALID_SESSION_USER_CODE,
			Message: common.INVALID_SESSION_USER_MESS,
		}, nil
	}

	user, err = l.svcCtx.UsersModel.FindOne(l.ctx, userId)
	if user == nil || err != nil {
		l.Logger.Error(common.DB_ERROR_CODE)
		return &types.GetUserRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	if userId != user.Id {
		l.Logger.Error(common.INVALID_SESSION_USER_CODE)
		return &types.GetUserRes{
			Code:    common.INVALID_SESSION_USER_CODE,
			Message: common.INVALID_SESSION_USER_MESS,
		}, nil
	}

	userOutPut := types.User{
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

	resp = &types.GetUserRes{
		Code:    common.SUCCESS_CODE,
		Message: common.SUCCESS_MESS,
		Data: types.GetUserData{
			User: userOutPut,
		},
	}
	l.Logger.Infof("Get User Succes: %v", resp)
	return
}
