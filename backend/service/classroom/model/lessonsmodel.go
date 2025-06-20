package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ LessonsModel = (*customLessonsModel)(nil)

type (
	// LessonsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLessonsModel.
	LessonsModel interface {
		lessonsModel
		withSession(session sqlx.Session) LessonsModel
	}

	customLessonsModel struct {
		*defaultLessonsModel
	}
)

// NewLessonsModel returns a model for the database table.
func NewLessonsModel(conn sqlx.SqlConn) LessonsModel {
	return &customLessonsModel{
		defaultLessonsModel: newLessonsModel(conn),
	}
}

func (m *customLessonsModel) withSession(session sqlx.Session) LessonsModel {
	return NewLessonsModel(sqlx.NewSqlConnFromSession(session))
}
