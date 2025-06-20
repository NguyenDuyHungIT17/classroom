package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ EmailConfirmationsModel = (*customEmailConfirmationsModel)(nil)

type (
	// EmailConfirmationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEmailConfirmationsModel.
	EmailConfirmationsModel interface {
		emailConfirmationsModel
		withSession(session sqlx.Session) EmailConfirmationsModel
	}

	customEmailConfirmationsModel struct {
		*defaultEmailConfirmationsModel
	}
)

// NewEmailConfirmationsModel returns a model for the database table.
func NewEmailConfirmationsModel(conn sqlx.SqlConn) EmailConfirmationsModel {
	return &customEmailConfirmationsModel{
		defaultEmailConfirmationsModel: newEmailConfirmationsModel(conn),
	}
}

func (m *customEmailConfirmationsModel) withSession(session sqlx.Session) EmailConfirmationsModel {
	return NewEmailConfirmationsModel(sqlx.NewSqlConnFromSession(session))
}
