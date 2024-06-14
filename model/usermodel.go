package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		withSession(session sqlx.Session) UserModel
		//添加新方法
		FindByUsername(ctx context.Context, username string) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (m *customUserModel) withSession(session sqlx.Session) UserModel {
	return NewUserModel(sqlx.NewSqlConnFromSession(session))
}

// 根据帐号获取用户信息
func (m *customUserModel) FindByUsername(ctx context.Context, username string) (*User, error) {
	var resp User

	//获取用户信息
	dataQuery := fmt.Sprintf("SELECT %s FROM %s WHERE delete_time = 0 AND username = ? LIMIT 1", userRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, dataQuery, username)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}
