package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"net/http"
	"usercenter/internal/svc"
	"usercenter/internal/types"
	"usercenter/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUserCreateLogic 新增用户
func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.UserCreateReq) (resp *types.UserCreateResp, err error) {
	logx.Info(l.ctx.Value("User-Agent"))
	if transErr := l.svcCtx.UserModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 预先声明响应体，确保始终返回
		resp = &types.UserCreateResp{
			Code:    http.StatusInternalServerError, // 默认设为错误状态
			Message: "创建用户失败",
			Flag:    false,
		}

		user := &model.User{
			Mobile:   req.Mobile,
			Nickname: req.Nickname,
		}
		// 添加 user
		dbRes, uErr := l.svcCtx.UserModel.TransInsert(ctx, nil, user)
		if uErr != nil {
			return uErr
		}
		userID, _ := dbRes.LastInsertId()

		// 添加UserData
		userData := &model.UserData{
			UserId: userID,
			Data:   "测试数据",
		}
		if _, udErr := l.svcCtx.UserDataModel.TransInsert(ctx, nil, userData); udErr != nil {
			return err
		}

		//return errors.New("手动触发失败.")
		return nil
	}); transErr == nil {
		return &types.UserCreateResp{
			Code:    http.StatusOK,
			Message: "创建用户成功",
			Flag:    true,
		}, nil
	} else {
		// 记录错误日志（可选）
		logx.WithContext(l.ctx).Errorf("创建用户失败: %v", transErr)
	}

	return resp, err
}
