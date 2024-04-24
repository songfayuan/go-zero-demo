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

type KafkaConsumerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKafkaConsumerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KafkaConsumerLogic {
	return &KafkaConsumerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KafkaConsumerLogic) KafkaConsumer() (resp *types.ApiKafkaResp, err error) {
	consumer, err := l.svcCtx.Sys.KafkaConsumer(l.ctx, &sysclient.Empty{})

	if err != nil {
		resJson, _ := json.Marshal(consumer)
		logx.WithContext(l.ctx).Errorf("Kafka消费者演示：操作失败，请求参数param = %s，异常信息errMsg = %s", resJson, err.Error())
		return nil, rpcerror.New(err)
	}

	return &types.ApiKafkaResp{
		Code:    200,
		Message: "操作成功",
		Data:    types.ApiKafkaReq{},
	}, nil
}
