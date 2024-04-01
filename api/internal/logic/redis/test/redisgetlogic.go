package test

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"go-zero-demo/common/errors/rpcerror"
	"go-zero-demo/rpc/sys/sysclient"

	"go-zero-demo/api/internal/svc"
	"go-zero-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedisGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisGetLogic {
	return &RedisGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedisGetLogic) RedisGet(req *types.ApiRedisGetReq) (resp *types.ApiRedisResp, err error) {
	param := &sysclient.RedisReq{}
	copier.Copy(param, req)
	getRes, err := l.svcCtx.Sys.RedisGet(l.ctx, param)

	if err != nil {
		resJson, _ := json.Marshal(getRes)
		logx.WithContext(l.ctx).Errorf("获取数据测试：操作失败，请求参数param = %s，异常信息errMsg = %s", resJson, err.Error())
		return nil, rpcerror.New(err)
	}

	return &types.ApiRedisResp{
		Code:    200,
		Message: "操作成功",
		Data: types.ApiRedisReq{
			Name:     getRes.Name,
			NickName: getRes.NickName,
			Password: getRes.Password,
			Email:    getRes.Email,
		},
	}, nil
}
