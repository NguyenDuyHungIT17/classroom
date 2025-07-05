package logic

import (
	"backend/common"
	"backend/service/classroom/api/internal/svc"
	"backend/service/classroom/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// DeleteUser
func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserReq) (resp *types.DeleteUserRes, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Infof("DeleteUser: %v", resp)

	err = l.svcCtx.UsersModel.Delete(l.ctx, req.UserId)
	if err != nil {
		l.Logger.Error(err)
		return &types.DeleteUserRes{
			Code:    common.DB_ERROR_CODE,
			Message: common.DB_ERROR_MESS,
		}, nil
	}

	resp = &types.DeleteUserRes{
		Code:    common.SUCCESS_CODE,
		Message: common.SUCCESS_MESS,
	}
	l.Logger.Infof("Delete User Success: %v ", resp)
	return resp, nil
}
