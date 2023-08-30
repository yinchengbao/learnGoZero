package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"login/internal/svc"
	"login/internal/types"
)

type MenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuLogic {
	return &MenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuLogic) MenuLogic(req *types.Request) (resp *types.Response, err error) {
	resp = &types.Response{
		Message: "进入菜单",
	}

	return
}
