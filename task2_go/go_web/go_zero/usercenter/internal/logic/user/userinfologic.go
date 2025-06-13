package user

import (
	"context"
	"errors"
	"net/http"
	"usercenter_rpc/usercenter_rpc"

	"usercenter/internal/svc"
	"usercenter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询用户信息
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &usercenter_rpc.GetUserInfoReq{
		Id: req.ID,
	})
	if err != nil {
		return &types.UserInfoResp{Code: http.StatusInternalServerError, Message: "查询失败"}, errors.New("查询失败")
	}
	resp = &types.UserInfoResp{Code: http.StatusOK, Message: "查询成功", Data: userResp.Nickname}
	return
}
