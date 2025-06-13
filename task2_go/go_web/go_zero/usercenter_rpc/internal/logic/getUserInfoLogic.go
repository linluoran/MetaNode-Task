package logic

import (
	"context"

	"usercenter_rpc/internal/svc"
	"usercenter_rpc/usercenter_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *usercenter_rpc.GetUserInfoReq) (*usercenter_rpc.GetUserInfoResp, error) {
	m := map[int64]string{
		1: "张三",
		2: "李四",
		3: "王五",
	}
	nickname := "unknown"
	if name, ok := m[in.Id]; ok {
		nickname = name
	}
	return &usercenter_rpc.GetUserInfoResp{
		Id:       in.Id,
		Nickname: nickname,
	}, nil
}
