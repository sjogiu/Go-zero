package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-study/api/model"

	"zero-study/api/demo/internal/svc"
	"zero-study/api/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建用户
func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {
	// todo: add your logic here and delete this line
	if err = l.svcCtx.UserModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := &model.User{}
		user.Mobile = req.Mobile
		user.Name = req.Name
		result, err := l.svcCtx.UserModel.TransInsert(ctx, session, user)
		if err != nil {
			return err
		}
		userId, _ := result.LastInsertId()
		userData := &model.UserData{
			Id:   userId,
			Data: "hxxx",
		}
		if _, err := l.svcCtx.UserDataModel.TransInsert(ctx, session, userData); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.New("创建用户失败")
	}
	return &types.CreateUserResp{
		Flag: true,
	}, nil
}
