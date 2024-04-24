package logic

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"

	"go-zero-demo/rpc/sys/internal/svc"
	"go-zero-demo/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type KafkaConsumerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKafkaConsumerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KafkaConsumerLogic {
	return &KafkaConsumerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Kafka消费者演示请求
// 这里演示手动请求触发kafka消费，实际项目中要做成项目启动后就一直监听消费。
func (l *KafkaConsumerLogic) KafkaConsumer(in *sysclient.Empty) (*sysclient.KafkaResp, error) {
	// 创建一个reader，指定GroupID，消费消息
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{l.svcCtx.KafkaConf.Host},
		GroupID:  l.svcCtx.KafkaConf.Group, //指定消费者组ID
		Topic:    l.svcCtx.KafkaConf.Topic,
		MaxBytes: 10e6, //10MB
	})

	//接收消息
	for {
		//ReadMessage会自动提交偏移量
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("Kafka消费者演示请求：message at topic/partition/offset %v/%v/%v: %s = %s\n", message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
	}

	//程序退出前关闭Reader
	if err := reader.Close(); err != nil {
		log.Fatal("Kafka消费者演示请求：failed to close reader:", err)
	}

	return &sysclient.KafkaResp{}, nil
}
