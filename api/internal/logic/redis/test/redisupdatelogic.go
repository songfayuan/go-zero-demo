package test

import (
	"context"
	"encoding/json"
	"go-zero-demo/common/errors/rpcerror"
	"go-zero-demo/rpc/sys/sysclient"

	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedisUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisUpdateLogic {
	return &RedisUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedisUpdateLogic) RedisUpdate(req *types.ApiRedisReq) (resp *types.ApiRedisResp, err error) {
	updateRes, err := l.svcCtx.Sys.RedisUpdate(l.ctx, &sysclient.RedisReq{
		Name:     req.Name,
		NickName: req.NickName,
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		resJson, _ := json.Marshal(updateRes)
		logx.WithContext(l.ctx).Errorf("Redis更新数据测试：操作失败，请求参数param = %s，异常信息errMsg = %s", resJson, err.Error())
		return nil, rpcerror.New(err)
	}

	return &types.ApiRedisResp{
		Code:    200,
		Message: "操作成功",
		Data: types.ApiRedisReq{
			Name:     updateRes.Name,
			NickName: updateRes.NickName,
			Password: updateRes.Password,
			Email:    updateRes.Email,
		},
	}, nil
}
