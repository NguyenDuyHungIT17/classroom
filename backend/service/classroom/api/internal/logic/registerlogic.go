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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Register
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	// todo: add your logic here and delete this line
	var user *model.Users
	currentTime := time.Now().UnixMilli()
	var passwordHash, token string
	var verificationCode, subject string
	var accessExpire = l.svcCtx.Config.Auth.AccessExpire
	var accessSecret = l.svcCtx.Config.Auth.AccessSecret
	//check email, phonenumber
	if !utils.ValidatesEmail(req.Email) || req.PhoneNumber != "" && !utils.ValidatesPhoneNumber(req.PhoneNumber) {
		l.Logger.Error(common.INVALID_REQUEST_MESS)
		return &types.RegisterRes{
			Code:    common.INVALID_REQUEST_CODE,
			Message: common.INVALID_REQUEST_MESS,
		}, nil
	}

	//exist user
	user, err = l.svcCtx.UsersModel.FindOneByUserNameOrEmail(l.ctx, req.UserName, req.Email)
	if err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	if user != nil {
		l.Logger.Error(common.USER_EXISTED_MESS)
		return &types.RegisterRes{
			Code:    common.USER_EXISTED_CODE,
			Message: common.USER_EXISTED_MESS,
		}, nil
	}

	passwordHash = utils.GetMD5Hasd(req.Password)
	verificationCode = utils.GenerateResetToken()

	user = &model.Users{
		UserName:         req.UserName,
		Password:         passwordHash,
		Email:            req.Email,
		PhoneNumber:      sql.NullString{String: req.PhoneNumber, Valid: true},
		Gender:           int64(req.Gender),
		FullName:         req.FullName,
		Avatar:           sql.NullString{Valid: false},
		IsVerified:       false,
		VerificationCode: sql.NullString{String: verificationCode, Valid: true},
		Role:             common.USER_ROLE_CUSTOMER,
		CreateTime:       currentTime,
		UpdateTime:       currentTime,
	}

	subject = "HDClassroom - Thông báo tài khoản"
	emailData := map[string]interface{}{
		"Subject":  subject,
		"UserName": user.UserName,
		"Password": req.Password,
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
	if err = utils.SendRegisterEmail(user.Email, utils.MailRegisterPath, smtpConfig, emailData); err != nil {
		l.Logger.Error(err)
		return &types.RegisterRes{
			Code:    common.SEND_EMAIL_ERROR_CODE,
			Message: common.SEND_EMAIL_ERROR_MESS,
		}, nil
	}

	//insert Db
	userId, err := l.svcCtx.UsersModel.InsertDb(l.ctx, user)
	if err != nil {
		return &types.RegisterRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	token, err = utils.GetJwtToken(accessSecret, currentTime, accessExpire, userId, int(user.Role))
	if err != nil {
		l.Logger.Error(common.INVALID_REQUEST_CODE)
		return &types.RegisterRes{
			Code:    common.INVALID_REQUEST_CODE,
			Message: common.INVALID_REQUEST_MESS,
		}, nil
	}

	resp = &types.RegisterRes{
		Code:    common.SUCCESS_CODE,
		Message: common.SUCCESS_MESS,
		Data: types.RegisterData{
			User: types.User{
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

	l.Logger.Infof("Register success: %v", resp)
	return
}
