package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
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
		FindMultipleByConditionWithPagging(ctx context.Context, mapConditions map[string]interface{}) (resp []*Users, err error)
		InsertDb(ctx context.Context, data *Users) (int64, error)
		UpdateDb(ctx context.Context, data *Users) error
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

func (m *customUsersModel) FindMultipleByConditionWithPagging(ctx context.Context, mapConditions map[string]interface{}) ([]*Users, error) {
	var args []interface{}
	var resp []*Users

	query := fmt.Sprintf("select %s from %s where 0 = 0", usersRows, m.table)

	if email, exist := mapConditions["email"].(string); exist && email != "" {
		query += " and `email` like `%" + email + "%`"
	}

	query += " order by `email`,`id` asc"

	if limit, exist := mapConditions["limit"].(int); exist && limit != 0 {
		query += " limit ?"
		args = append(args, limit)
	}

	if offset, exist := mapConditions["offset"].(int); exist && offset != 0 {
		query += " offset ?"
		args = append(args, offset)
	}

	logx.Info(query)
	logx.Info(args...)

	err := m.conn.QueryRowsCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUsersModel) InsertDb(ctx context.Context, data *Users) (int64, error) {
	var usersRows = "`user_name`, `password`, `email`, `phone_number`, `gender`, `full_name`, `avatar`, `is_verified`, `verification_code`, `role`, `create_time`, `update_time`"
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

func (m *customUsersModel) UpdateDb(ctx context.Context, data *Users) error {
	query := fmt.Sprintf("update %s set `user_name` = ?, `password` = ?, `email` = ?, `phone_number` = ?, `gender` = ?, `full_name` = ?, `avatar` = ?, `is_verified` = ?, `verification_code` = ?, `role` = ?, `create_time` = ?, `update_time` = ? where `id` = ?", m.table)

	_, err := m.conn.ExecCtx(ctx, query,
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
		data.Id,
	)

	return err
}
