package user

import (
	"context"
	"database/sql"
	"errors"

	"demo/internal/svc"
	"demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 演示模块：根据名称返回问候语
func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoLogic) Demo(req *types.Request) (resp *types.Response, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserID)
	if errors.Is(err, sql.ErrNoRows) {
		return &types.Response{Message: "用户不存在"}, errors.New("用户不存在")
	}

	if err != nil {
		return &types.Response{Message: "查询失败"}, errors.New("查询失败")
	}

	resp = &types.Response{Message: user.Nickname}
	return
}
