package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ClassesModel = (*customClassesModel)(nil)

type (
	// ClassesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClassesModel.
	ClassesModel interface {
		classesModel
		withSession(session sqlx.Session) ClassesModel
		InsertDb(ctx context.Context, data *Classes) (int64, error)
		ExistsByClassCode(ctx context.Context, classCode string) (bool, error)
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

func (m *customClassesModel) ExistsByClassCode(ctx context.Context, classCode string) (bool, error) {
	var count int64
	query := fmt.Sprintf("select count(*) from %s where `class_code` = ?", m.table)
	err := m.conn.QueryRowCtx(ctx, &count, query, classCode)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (m *customClassesModel) InsertDb(ctx context.Context, data *Classes) (int64, error) {
	classRows := "`class_code`, `class_name`, `description`, `teacher_id`, `create_time`, `update_time`"
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, classRows)
	result, err := m.conn.ExecCtx(ctx, query,
		data.ClassCode,
		data.ClassName,
		data.Description,
		data.TeacherId,
		data.CreateTime,
		data.UpdateTime,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
