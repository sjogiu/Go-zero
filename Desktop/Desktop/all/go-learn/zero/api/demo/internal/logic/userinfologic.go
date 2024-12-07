package logic

import (
	"context"
	"errors"

	"zero-study/api/demo/internal/svc"
	"zero-study/api/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	// todo: add your logic here and delete this line
	logx.Info("info level")
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errors.New("查询出错")
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	resp = &types.GetUserInfoResp{
		Id:   user.Id,
		Name: user.Name,
	}
	return
}
