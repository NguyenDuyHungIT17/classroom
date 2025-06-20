package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GradeComponentsModel = (*customGradeComponentsModel)(nil)

type (
	// GradeComponentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGradeComponentsModel.
	GradeComponentsModel interface {
		gradeComponentsModel
		withSession(session sqlx.Session) GradeComponentsModel
	}

	customGradeComponentsModel struct {
		*defaultGradeComponentsModel
	}
)

// NewGradeComponentsModel returns a model for the database table.
func NewGradeComponentsModel(conn sqlx.SqlConn) GradeComponentsModel {
	return &customGradeComponentsModel{
		defaultGradeComponentsModel: newGradeComponentsModel(conn),
	}
}

func (m *customGradeComponentsModel) withSession(session sqlx.Session) GradeComponentsModel {
	return NewGradeComponentsModel(sqlx.NewSqlConnFromSession(session))
}
