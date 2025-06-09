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
		TransUpdate(ctx context.Context, session sqlx.Session, data *User) error
		TransDelete(ctx context.Context, session sqlx.Session, id int64) error
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

// TransInsert Insert事务
func (m *customUserModel) TransInsert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	gozeroUserIdKey := fmt.Sprintf("%s%v", cacheGozeroUserIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Nickname, data.Mobile)
		}
		return conn.ExecCtx(ctx, query, data.Nickname, data.Mobile)
	}, gozeroUserIdKey)
	return ret, err
}

// TransDelete Delete事务
func (m *customUserModel) TransDelete(ctx context.Context, session sqlx.Session, id int64) error {
	gozeroUserIdKey := fmt.Sprintf("%s%v", cacheGozeroUserIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, gozeroUserIdKey)
	return err
}

// TransUpdate 更新事务
func (m *customUserModel) TransUpdate(ctx context.Context, session sqlx.Session, data *User) error {
	gozeroUserIdKey := fmt.Sprintf("%s%v", cacheGozeroUserIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Nickname, data.Mobile, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.Nickname, data.Mobile, data.Id)
	}, gozeroUserIdKey)
	return err
}
