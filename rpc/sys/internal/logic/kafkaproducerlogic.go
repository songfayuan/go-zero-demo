package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"

	"go-zero-demo/rpc/sys/internal/svc"
	"go-zero-demo/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type KafkaProducerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKafkaProducerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KafkaProducerLogic {
	return &KafkaProducerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Kafka生产者演示请求
func (l *KafkaProducerLogic) KafkaProducer(in *sysclient.KafkaReq) (*sysclient.KafkaResp, error) {
	if in.Name == "" {
		return nil, errors.New("账号不能为空")
	}
	if in.NickName == "" {
		return nil, errors.New("姓名不能为空")
	}
	if in.Email == "" {
		return nil, errors.New("邮箱不能为空")
	}

	// 创建一个writer，向topic发送消息
	w := &kafka.Writer{
		Addr:         kafka.TCP(l.svcCtx.Config.KafkaConf.Host),
		Topic:        l.svcCtx.Config.KafkaConf.Topic,
		Balancer:     &kafka.LeastBytes{}, // 指定分区的balancer模式为最小字节分布
		RequiredAcks: kafka.RequireAll,    // ack模式
		Async:        true,
	}

	// 定义消息内容
	messages := []string{in.Name, in.NickName, in.Password, in.Email}

	// 循环发送消息
	for i, msg := range messages {
		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("Key-%d", i+1)), // 使用不同的分区键
				Value: []byte(msg),
			},
		)
		if err != nil {
			log.Fatalf("Kafka生产者演示请求，向kafka写入数据失败: %v", err)
		}
	}

	if err := w.Close(); err != nil {
		log.Fatal("Kafka生产者演示请求，failed to close writer:", err)
	}

	return &sysclient.KafkaResp{}, nil
}
