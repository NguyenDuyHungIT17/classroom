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

		InsertDb(ctx context.Context, data *Enrollments) error
		DeleteByClassId(ctx context.Context, classId int64) error
		ExistsByClassIdAndStudentId(ctx context.Context, classId, studentId int64) (bool, error)
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

func (m *customEnrollmentsModel) ExistsByClassIdAndStudentId(ctx context.Context, classId, studentId int64) (bool, error) {
	var count int64
	query := fmt.Sprintf("select count(*) from %s where `student_id` = ? and `class_id` = ?", m.table)
	err := m.conn.QueryRowCtx(ctx, &count, query, studentId, classId)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (m *customEnrollmentsModel) InsertDb(ctx context.Context, data *Enrollments) error {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, enrollmentsRowsExpectAutoSet)
	_, err := m.conn.ExecCtx(ctx, query, data.StudentId, data.ClassId, data.JoinTime)
	return err
}
