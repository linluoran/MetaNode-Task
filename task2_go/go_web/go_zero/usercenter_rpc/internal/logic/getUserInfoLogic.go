package logic

import (
	"context"

	"usercenter_rpc/internal/svc"
	"usercenter_rpc/usercenter_rpc"

	"github.com/jinzhu/copier"
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

	var userModelTmp usercenter_rpc.GetUserInfoResp
	user, _ := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)

	// 这个包博主推荐 前提是类型和字段名必须一致
	err := copier.Copy(&userModelTmp, user)
	if err != nil {
		return nil, err
	}

	return &userModelTmp, nil
}
