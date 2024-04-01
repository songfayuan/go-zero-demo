package test

import (
	"context"
	"go-zero-demo/rpc/sys/sysclient"

	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedisDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisDeleteLogic {
	return &RedisDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedisDeleteLogic) RedisDelete(req *types.ApiRedisReq) (resp *types.ApiRedisResp, err error) {
	l.svcCtx.Sys.RedisDelete(l.ctx, &sysclient.RedisReq{})
	return &types.ApiRedisResp{
		Code:    200,
		Message: "操作成功",
		Data:    types.ApiRedisReq{},
	}, nil
}
