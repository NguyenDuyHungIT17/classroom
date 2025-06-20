package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		withSession(session sqlx.Session) UsersModel
		FindOneByUserNameOrEmail(ctx context.Context, username, email string) (*Users, error)
		InsertDb(ctx context.Context, data *Users) (int64, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn),
	}
}

func (m *customUsersModel) withSession(session sqlx.Session) UsersModel {
	return NewUsersModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customUsersModel) FindOneByUserNameOrEmail(ctx context.Context, username, email string) (*Users, error) {
	var resp Users

	query := fmt.Sprintf("select %s from %s where `email` = ? or `user_name` = ? limit 1", usersRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, email, username)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customUsersModel) InsertDb(ctx context.Context, data *Users) (int64, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usersRows)
	result, err := m.conn.ExecCtx(ctx, query,
		data.UserName,
		data.Password,
		data.Email,
		data.PhoneNumber,
		data.Gender,
		data.FullName,
		data.Avatar,
		data.IsVerified,
		data.VerificationCode,
		data.Role,
		data.CreateTime,
		data.UpdateTime,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
