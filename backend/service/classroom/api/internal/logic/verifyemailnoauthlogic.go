package logic

import (
	"context"
	"database/sql"
	"time"

	"backend/common"
	"backend/service/classroom/api/internal/svc"
	"backend/service/classroom/api/internal/types"
	"backend/service/classroom/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyEmailNoAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// VerifyEmailNoAuth
func NewVerifyEmailNoAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyEmailNoAuthLogic {
	return &VerifyEmailNoAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyEmailNoAuthLogic) VerifyEmailNoAuth(req *types.VerifyEmailNoAuthReq) (resp *types.VerifyEmailNoAuthRes, err error) {
	// todo: add your logic here and delete this line
	var user *model.Users
	currentTime := time.Now().UnixMilli()

	user, err = l.svcCtx.UsersModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil || user == nil {
		l.Logger.Error(err)
		return &types.VerifyEmailNoAuthRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	if user.VerificationCode.String != req.Token {
		l.Logger.Error(common.INVALID_TOKEN_MESS)
		return &types.VerifyEmailNoAuthRes{
			Code:    common.INVALID_TOKEN_CODE,
			Message: common.INVALID_TOKEN_MESS,
		}, nil
	}

	user.VerificationCode = sql.NullString{String: "", Valid: true}
	user.UpdateTime = currentTime
	if err = l.svcCtx.UsersModel.UpdateDb(l.ctx, user); err != nil {
		l.Logger.Error(err)
		return &types.VerifyEmailNoAuthRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}
	resp = &types.VerifyEmailNoAuthRes{
		Code:    common.SUCCESS_CODE,
		Message: common.SUCCESS_MESS,
	}

	l.Logger.Infof("VerifyEmailNoAuth success: %v", resp)
	return resp, nil
}
