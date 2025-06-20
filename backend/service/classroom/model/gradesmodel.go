package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GradesModel = (*customGradesModel)(nil)

type (
	// GradesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGradesModel.
	GradesModel interface {
		gradesModel
		withSession(session sqlx.Session) GradesModel
	}

	customGradesModel struct {
		*defaultGradesModel
	}
)

// NewGradesModel returns a model for the database table.
func NewGradesModel(conn sqlx.SqlConn) GradesModel {
	return &customGradesModel{
		defaultGradesModel: newGradesModel(conn),
	}
}

func (m *customGradesModel) withSession(session sqlx.Session) GradesModel {
	return NewGradesModel(sqlx.NewSqlConnFromSession(session))
}
