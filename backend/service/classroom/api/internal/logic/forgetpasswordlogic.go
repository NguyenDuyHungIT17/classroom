package logic

import (
	"context"
	"database/sql"
	"time"

	"backend/common"
	"backend/service/classroom/api/internal/svc"
	"backend/service/classroom/api/internal/types"
	"backend/service/classroom/model"
	"backend/service/classroom/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForgetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// ForgetPassword
func NewForgetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgetPasswordLogic {
	return &ForgetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ForgetPasswordLogic) ForgetPassword(req *types.ForgetPasswordReq) (resp *types.ForgetPasswordRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Infof("Forget password: %v", req)

	var user *model.Users
	var token, url, subject string
	var currentTime = time.Now().UnixMilli()

	if req.UserName != "" || req.Email != "" {
		user, err = l.svcCtx.UsersModel.FindOneByUserNameOrEmail(l.ctx, req.UserName, req.Email)
		if err != nil {
			l.Logger.Error(err)
			return &types.ForgetPasswordRes{
				Code:    common.DB_ERROR_CODE,
				Message: common.DB_ERROR_MESS,
			}, nil
		}
	} else {
		l.Logger.Error(common.INVALID_REQUEST_MESS)
		return &types.ForgetPasswordRes{
			Code:    common.INVALID_REQUEST_CODE,
			Message: common.INVALID_REQUEST_MESS,
		}, nil
	}

	if user == nil {
		l.Logger.Error(common.USER_IS_NOT_EXIST_MESS)
		return &types.ForgetPasswordRes{
			Code:    common.USER_IS_NOT_EXIST_CODE,
			Message: common.USER_IS_NOT_EXIST_MESS,
		}, nil
	}

	token = utils.GenerateResetToken()
	url = l.svcCtx.Config.SMTPConfig.ClientOrigin + "/verify-email?email=" + user.Email + "&token=" + token
	subject = "HDManager - Quên mật khẩu"
	emailData := utils.EmailData{
		URL:      url,
		Subject:  subject,
		UserName: req.UserName,
	}
	smtpConfig := utils.SMTPConfig{
		EmailFrom:    l.svcCtx.Config.SMTPConfig.EmailFrom,
		SMTPHost:     l.svcCtx.Config.SMTPConfig.SMTPHost,
		SMTPPass:     l.svcCtx.Config.SMTPConfig.SMTPPass,
		SMTPPort:     int(l.svcCtx.Config.SMTPConfig.SMTPPort),
		SMTPUser:     l.svcCtx.Config.SMTPConfig.SMTPUser,
		ClientOrigin: l.svcCtx.Config.SMTPConfig.ClientOrigin,
	}
	//send email
	if err = utils.SendEmail(user.Email, utils.MailResetPassword, smtpConfig, emailData); err != nil {
		l.Logger.Error(err)
		return &types.ForgetPasswordRes{
			Code:    common.SEND_EMAIL_ERROR_CODE,
			Message: common.SEND_EMAIL_ERROR_MESS,
		}, nil
	}

	user.VerificationCode = sql.NullString{String: token, Valid: true}
	user.UpdateTime = currentTime

	if err = l.svcCtx.UsersModel.UpdateDb(l.ctx, user); err != nil {
		l.Logger.Error(err)
		return &types.ForgetPasswordRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	resp = &types.ForgetPasswordRes{
		Code:    common.SUCCESS_CODE,
		Message: common.SUCCESS_MESS,
	}

	l.Logger.Infof("ForgetPassword success: %v", resp)
	return resp, nil
}
