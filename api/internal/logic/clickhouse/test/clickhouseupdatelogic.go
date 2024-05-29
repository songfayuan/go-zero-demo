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

type ClickhouseUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClickhouseUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClickhouseUpdateLogic {
	return &ClickhouseUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClickhouseUpdateLogic) ClickhouseUpdate(req *types.ApiClickhouseReq) (resp *types.ApiClickhouseResp, err error) {
	updateRes, err := l.svcCtx.Sys.ClickhouseUpdate(l.ctx, &sysclient.ClickhouseReq{
		Name:     req.Name,
		NickName: req.NickName,
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		resJson, _ := json.Marshal(updateRes)
		logx.WithContext(l.ctx).Errorf("Clickhouse更新数据测试：操作失败，请求参数param = %s，异常信息errMsg = %s", resJson, err.Error())
		return nil, rpcerror.New(err)
	}

	return &types.ApiClickhouseResp{
		Code:    200,
		Message: "操作成功",
		Data: types.ApiClickhouseReq{
			Name:     updateRes.Name,
			NickName: updateRes.NickName,
			Password: updateRes.Password,
			Email:    updateRes.Email,
		},
	}, nil
}
