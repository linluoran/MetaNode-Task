package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserDataModel = (*customUserDataModel)(nil)

type (
	// UserDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserDataModel.
	UserDataModel interface {
		userDataModel
		// TransCtx 添加事务方法
		TransCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error
	}

	customUserDataModel struct {
		*defaultUserDataModel
	}
)

// NewUserDataModel returns a model for the database table.
func NewUserDataModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserDataModel {
	return &customUserDataModel{
		defaultUserDataModel: newUserDataModel(conn, c, opts...),
	}
}

// TransCtx 实现事务方法
func (m *customUserDataModel) TransCtx(ctx context.Context, fn func(context.Context, sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}
