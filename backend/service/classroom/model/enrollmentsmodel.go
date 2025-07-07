package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ EnrollmentsModel = (*customEnrollmentsModel)(nil)

type (
	// EnrollmentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEnrollmentsModel.
	EnrollmentsModel interface {
		enrollmentsModel
		withSession(session sqlx.Session) EnrollmentsModel

		DeleteByClassId(ctx context.Context, classId int64) error
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

func (m *customEnrollmentsModel) DeleteByClassId(ctx context.Context, classId int64) error {
	query := fmt.Sprintf("delete from %s where `class_id` = ? ", m.table)
	_, err := m.conn.ExecCtx(ctx, query, classId)
	return err
}
