package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SubjectsModel = (*customSubjectsModel)(nil)

type (
	// SubjectsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSubjectsModel.
	SubjectsModel interface {
		subjectsModel
		withSession(session sqlx.Session) SubjectsModel
	}

	customSubjectsModel struct {
		*defaultSubjectsModel
	}
)

// NewSubjectsModel returns a model for the database table.
func NewSubjectsModel(conn sqlx.SqlConn) SubjectsModel {
	return &customSubjectsModel{
		defaultSubjectsModel: newSubjectsModel(conn),
	}
}

func (m *customSubjectsModel) withSession(session sqlx.Session) SubjectsModel {
	return NewSubjectsModel(sqlx.NewSqlConnFromSession(session))
}
