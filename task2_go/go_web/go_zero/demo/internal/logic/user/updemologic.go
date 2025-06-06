package user

import (
	"context"

	"demo/internal/svc"
	"demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UPDemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUPDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UPDemoLogic {
	return &UPDemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UPDemoLogic) UPDemo(req *types.UpRequest) (resp *types.UpResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
