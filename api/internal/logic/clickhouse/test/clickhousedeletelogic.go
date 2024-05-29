package test

import (
	"context"
	"go-zero-demo/rpc/sys/sysclient"

	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClickhouseDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClickhouseDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClickhouseDeleteLogic {
	return &ClickhouseDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClickhouseDeleteLogic) ClickhouseDelete(req *types.ApiClickhouseReq) (resp *types.ApiClickhouseResp, err error) {
	l.svcCtx.Sys.ClickhouseDelete(l.ctx, &sysclient.ClickhouseReq{})
	return &types.ApiClickhouseResp{
		Code:    200,
		Message: "删除成功",
		Data:    types.ApiClickhouseReq{},
	}, nil
}
