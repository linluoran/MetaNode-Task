package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		// TransCtx 添加事务方法
		TransCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error
		TransInsert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}

// TransCtx 实现事务方法
func (m *customUserModel) TransCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

// TransInsert 事务
func (m *defaultUserModel) TransInsert(ctx context.Context, serssion sqlx.Session, data *User) (sql.Result, error) {
	gozeroUserIdKey := fmt.Sprintf("%s%v", cacheGozeroUserIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userRowsExpectAutoSet)
		return serssion.ExecCtx(ctx, query, data.Nickname, data.Mobile)
	}, gozeroUserIdKey)
	return ret, err
}
