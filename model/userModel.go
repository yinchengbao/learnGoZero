package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"login/internal/types"
)

type (
	UserModel interface {
		FindALL(ctx context.Context) ([]*types.User, error)
	}

	defaultUserModel struct {
		conn  sqlx.SqlConn
		table string
	}
)

func NewUserModel(coon sqlx.SqlConn) *defaultUserModel {
	return &defaultUserModel{
		conn:  coon,
		table: "`users`",
	}
}
func (m *defaultUserModel) FindALL(ctx context.Context) ([]*types.User, error) {
	sql := "select * from users"
	var userList []*types.User
	err := m.conn.QueryRowsCtx(ctx, &userList, sql)
	return userList, err
}
