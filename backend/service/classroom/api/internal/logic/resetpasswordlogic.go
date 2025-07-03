package logic

import (
	"context"
	"encoding/json"
	"time"

	"backend/common"
	"backend/service/classroom/api/internal/svc"
	"backend/service/classroom/api/internal/types"
	"backend/service/classroom/model"
	"backend/service/classroom/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// ResetPassword
func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordLogic) ResetPassword(req *types.ResetPasswordReq) (resp *types.ResetPasswordRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Infof("ForgetPassword: %v", req)

	var user *model.Users
	var userId int64
	currentTime := time.Now().UnixMilli()

	userId, err = l.ctx.Value("userId").(json.Number).Int64()
	if userId == 0 || err != nil {
		l.Logger.Error(err)
		return &types.ResetPasswordRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	user, err = l.svcCtx.UsersModel.FindOneByUserName(l.ctx, req.Username)
	if user == nil || err != nil {
		l.Logger.Error(err)
		return &types.ResetPasswordRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	if userId != user.Id {
		l.Logger.Error(common.INVALID_SESSION_USER_MESS)
		return &types.ResetPasswordRes{
			Code:    common.INVALID_SESSION_USER_CODE,
			Message: common.INVALID_SESSION_USER_MESS,
		}, nil
	}

	if user.Password == utils.GetMD5Hasd(req.PassWord) {
		l.Logger.Error(common.INVALID_SESSION_USER_MESS)
		return &types.ResetPasswordRes{
			Code:    common.INVALID_SESSION_USER_CODE,
			Message: common.INVALID_SESSION_USER_MESS,
		}, nil
	}

	user.Password = utils.GetMD5Hasd(req.PassWord)
	user.UpdateTime = currentTime
	if err = l.svcCtx.UsersModel.UpdateDb(l.ctx, user); err != nil {
		l.Logger.Error(err)
		return &types.ResetPasswordRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	resp = &types.ResetPasswordRes{
		Code:    common.SUCCESS_CODE,
		Message: common.SUCCESS_MESS,
	}

	l.Logger.Infof("ForgetPassword: %v", resp)
	return resp, nil
}
