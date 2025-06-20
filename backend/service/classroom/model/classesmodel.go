package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ClassesModel = (*customClassesModel)(nil)

type (
	// ClassesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClassesModel.
	ClassesModel interface {
		classesModel
		withSession(session sqlx.Session) ClassesModel
	}

	customClassesModel struct {
		*defaultClassesModel
	}
)

// NewClassesModel returns a model for the database table.
func NewClassesModel(conn sqlx.SqlConn) ClassesModel {
	return &customClassesModel{
		defaultClassesModel: newClassesModel(conn),
	}
}

func (m *customClassesModel) withSession(session sqlx.Session) ClassesModel {
	return NewClassesModel(sqlx.NewSqlConnFromSession(session))
}
