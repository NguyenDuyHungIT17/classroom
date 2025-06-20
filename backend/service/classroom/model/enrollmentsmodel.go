package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ EnrollmentsModel = (*customEnrollmentsModel)(nil)

type (
	// EnrollmentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEnrollmentsModel.
	EnrollmentsModel interface {
		enrollmentsModel
		withSession(session sqlx.Session) EnrollmentsModel
	}

	customEnrollmentsModel struct {
		*defaultEnrollmentsModel
	}
)

// NewEnrollmentsModel returns a model for the database table.
func NewEnrollmentsModel(conn sqlx.SqlConn) EnrollmentsModel {
	return &customEnrollmentsModel{
		defaultEnrollmentsModel: newEnrollmentsModel(conn),
	}
}

func (m *customEnrollmentsModel) withSession(session sqlx.Session) EnrollmentsModel {
	return NewEnrollmentsModel(sqlx.NewSqlConnFromSession(session))
}
