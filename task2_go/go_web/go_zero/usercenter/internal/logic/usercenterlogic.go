package logic

import (
	"context"

	"go_zero/usercenter/internal/svc"
	"go_zero/usercenter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsercenterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsercenterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsercenterLogic {
	return &UsercenterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsercenterLogic) Usercenter(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
