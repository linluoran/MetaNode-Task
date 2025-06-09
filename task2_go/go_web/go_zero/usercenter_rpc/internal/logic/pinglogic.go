package logic

import (
	"context"

	"usercenter_rpc/internal/svc"
	"usercenter_rpc/usercenter_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *usercenter_rpc.Request) (*usercenter_rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &usercenter_rpc.Response{}, nil
}
