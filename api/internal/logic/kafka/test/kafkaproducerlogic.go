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

type KafkaProducerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKafkaProducerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KafkaProducerLogic {
	return &KafkaProducerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KafkaProducerLogic) KafkaProducer(req *types.ApiKafkaReq) (resp *types.ApiKafkaResp, err error) {
	producer, err := l.svcCtx.Sys.KafkaProducer(l.ctx, &sysclient.KafkaReq{
		Name:     req.Name,
		NickName: req.NickName,
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		resJson, _ := json.Marshal(producer)
		logx.WithContext(l.ctx).Errorf("Kafka生产者演示：操作失败，请求参数param = %s，异常信息errMsg = %s", resJson, err.Error())
		return nil, rpcerror.New(err)
	}

	return &types.ApiKafkaResp{
		Code:    200,
		Message: "操作成功",
		Data: types.ApiKafkaReq{
			Name:     producer.Name,
			NickName: producer.NickName,
			Password: producer.Password,
			Email:    producer.Email,
		},
	}, nil
}
