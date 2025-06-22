package logic

import (
	"context"
	"time"

	"backend/common"
	"backend/service/classroom/api/internal/svc"
	"backend/service/classroom/api/internal/types"
	"backend/service/classroom/model"
	"backend/service/classroom/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Login
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("Login: %v", req)
	var user *model.Users
	var token string
	var accessExpire = l.svcCtx.Config.Auth.AccessExpire
	var accessSecret = l.svcCtx.Config.Auth.AccessSecret
	currentTime := time.Now().Unix()
	user, err = l.svcCtx.UsersModel.FindOneByUserName(l.ctx, req.UserName)
	if err != nil {
		l.Logger.Error(err)
		return &types.LoginRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	if user == nil {
		l.Logger.Error(common.USER_IS_NOT_EXIST_MESS)
		return &types.LoginRes{
			Code:    common.USER_IS_NOT_EXIST_CODE,
			Message: common.USER_IS_NOT_EXIST_MESS,
		}, nil
	}

	if utils.GetMD5Hasd(req.PassWord) != user.Password {
		l.Logger.Error(common.PASSWORD_IS_WRONG_MESS)
		return &types.LoginRes{
			Code:    common.PASSWORD_IS_WRONG_CODE,
			Message: common.PASSWORD_IS_WRONG_MESS,
		}, nil
	}

	//sinh token
	token, err = utils.GetJwtToken(accessSecret, currentTime, accessExpire, user.Id, int(user.Role))
	if err != nil {
		l.Logger.Error(common.INVALID_REQUEST_MESS)
		return &types.LoginRes{
			Code:    common.INVALID_REQUEST_CODE,
			Message: common.INVALID_REQUEST_MESS,
		}, nil
	}

	resp = &types.LoginRes{
		Code:    common.SUCCESS_CODE,
		Message: common.SUCCESS_MESS,
		Data: types.LoginData{
			User: types.User{
				Id:          user.Id,
				UserName:    user.UserName,
				Email:       user.Email,
				PhoneNumber: user.PhoneNumber.String,
				Gender:      int(user.Gender),
				FullName:    user.FullName,
				Avatar:      "",
				IsVerified:  user.IsVerified,
				Role:        int(user.Role),
				CreateTime:  user.CreateTime,
				UpdateTime:  user.UpdateTime,
			},
			Token: token,
		},
	}
	l.Logger.Infof("Login success: %v", resp)
	return
}
