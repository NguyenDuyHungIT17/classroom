package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TeachersModel = (*customTeachersModel)(nil)

type (
	// TeachersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTeachersModel.
	TeachersModel interface {
		teachersModel
		withSession(session sqlx.Session) TeachersModel

		InsertDb(ctx context.Context, data *Teachers) error
		DeleteByClassId(ctx context.Context, classId int64) error
	}

	customTeachersModel struct {
		*defaultTeachersModel
	}
)

// NewTeachersModel returns a model for the database table.
func NewTeachersModel(conn sqlx.SqlConn) TeachersModel {
	return &customTeachersModel{
		defaultTeachersModel: newTeachersModel(conn),
	}
}

func (m *customTeachersModel) withSession(session sqlx.Session) TeachersModel {
	return NewTeachersModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTeachersModel) InsertDb(ctx context.Context, data *Teachers) error {
	teacherRows := fmt.Sprintf("`user_id`, `bio`, `create_time`, `update_time`, `class_id`")
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, teacherRows)
	_, err := m.conn.ExecCtx(ctx, query,
		data.UserId,
		data.Bio,
		data.CreateTime,
		data.UpdateTime,
		data.ClassId,
	)
	return err
}

func (m *customTeachersModel) DeleteByClassId(ctx context.Context, classId int64) error {
	query := fmt.Sprintf("delete from %s where `class_id` = ? ", m.table)
	_, err := m.conn.ExecCtx(ctx, query, classId)
	return err
}
