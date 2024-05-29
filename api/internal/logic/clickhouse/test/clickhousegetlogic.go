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

type ClickhouseGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClickhouseGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClickhouseGetLogic {
	return &ClickhouseGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClickhouseGetLogic) ClickhouseGet(req *types.ApiClickhouseGetReq) (resp *types.ApiClickhouseResp, err error) {
	param := &sysclient.ClickhouseReq{}
	copier.Copy(param, req)
	getRes, err := l.svcCtx.Sys.ClickhouseGet(l.ctx, param)

	if err != nil {
		resJson, _ := json.Marshal(getRes)
		logx.WithContext(l.ctx).Errorf("获取数据测试：操作失败，请求参数param = %s，异常信息errMsg = %s", resJson, err.Error())
		return nil, rpcerror.New(err)
	}

	return &types.ApiClickhouseResp{
		Code:    200,
		Message: "操作成功",
		Data: types.ApiClickhouseReq{
			Name:     getRes.Name,
			NickName: getRes.NickName,
			Password: getRes.Password,
			Email:    getRes.Email,
		},
	}, nil
}
