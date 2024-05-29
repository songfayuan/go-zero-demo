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

type ClickhouseAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClickhouseAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClickhouseAddLogic {
	return &ClickhouseAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClickhouseAddLogic) ClickhouseAdd(req *types.ApiClickhouseReq) (resp *types.ApiClickhouseResp, err error) {
	res, err := l.svcCtx.Sys.ClickhouseAdd(l.ctx, &sysclient.ClickhouseReq{
		Name:     req.Name,
		NickName: req.NickName,
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		reqJson, _ := json.Marshal(res)
		logx.WithContext(l.ctx).Errorf("新增Clickhouse信息失败，请求参数：%s，异常信息：%s", reqJson, err.Error())
		return nil, rpcerror.New(err)
	}

	return &types.ApiClickhouseResp{
		Code:    200,
		Message: "新增成功",
		Data: types.ApiClickhouseReq{
			Name:     req.Name,
			NickName: req.NickName,
			Password: req.Password,
			Email:    req.Email,
		},
	}, nil
}
