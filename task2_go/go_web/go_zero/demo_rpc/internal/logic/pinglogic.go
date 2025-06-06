package logic

import (
	"context"

	"demo_rpc/demo_rpc"
	"demo_rpc/internal/svc"

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

func (l *PingLogic) Ping(in *demo_rpc.Request) (*demo_rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &demo_rpc.Response{}, nil
}
