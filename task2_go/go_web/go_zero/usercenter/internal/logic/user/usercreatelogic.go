package user

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_zero/usercenter/model"

	"go_zero/usercenter/internal/svc"
	"go_zero/usercenter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增用户
func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.UserCreateReq) (resp *types.UserCreateResp, err error) {
	if err := l.svcCtx.UserModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := &model.User{
			Mobile:   req.Mobile,
			Nickname: req.Nickname,
		}
		// 添加 user
		dbRes, err := l.svcCtx.UserModel.T
		if err != nil {
			return err
		}

		userID, _ := dbRes.LastInsertId()

		// 添加UserData
		userData := &model.UserData{
			Id:   userID,
			Data: "xxx",
		}
		if _, err := l.svcCtx.UserDataModel.Insert(ctx, userData); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.New("create user failed")
	}

	return
}
