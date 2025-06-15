package model

import (
	"context"
	"database/sql"
	"fmt"
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
		TransInsert(ctx context.Context, session sqlx.Session, data *UserData) (sql.Result, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *UserData) error
		TransDelete(ctx context.Context, session sqlx.Session, id int64) error
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

// TransDelete Delete事务
func (m *customUserDataModel) TransDelete(ctx context.Context, session sqlx.Session, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	gozeroUserDataIdKey := fmt.Sprintf("%s%v", cacheGozeroUserDataIdPrefix, id)
	gozeroUserDataUserIdKey := fmt.Sprintf("%s%v", cacheGozeroUserDataUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, gozeroUserDataIdKey, gozeroUserDataUserIdKey)
	return err
}

// TransInsert Insert事务
func (m *customUserDataModel) TransInsert(ctx context.Context, session sqlx.Session, data *UserData) (sql.Result, error) {
	gozeroUserDataIdKey := fmt.Sprintf("%s%v", cacheGozeroUserDataIdPrefix, data.Id)
	gozeroUserDataUserIdKey := fmt.Sprintf("%s%v", cacheGozeroUserDataUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userDataRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.UserId, data.Data)
		}
		return conn.ExecCtx(ctx, query, data.UserId, data.Data)
	}, gozeroUserDataIdKey, gozeroUserDataUserIdKey)
	return ret, err
}

// TransUpdate 更新事务
func (m *customUserDataModel) TransUpdate(ctx context.Context, session sqlx.Session, newData *UserData) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	gozeroUserDataIdKey := fmt.Sprintf("%s%v", cacheGozeroUserDataIdPrefix, data.Id)
	gozeroUserDataUserIdKey := fmt.Sprintf("%s%v", cacheGozeroUserDataUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userDataRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, newData.UserId, newData.Data, newData.Id)
		}
		return conn.ExecCtx(ctx, query, newData.UserId, newData.Data, newData.Id)
	}, gozeroUserDataIdKey, gozeroUserDataUserIdKey)
	return err
}
