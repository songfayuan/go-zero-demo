package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

// 演示kafka读写
func main() {
	//writeByConn()
	writeByWriter3()
	//writeByWriter2()
	//readByConn()
	//readByReader()
	//readByReaderGroup()
}

// writeByConn基于Conn发送消息
func writeByConn() {
	topic := "kafka-test-topic3"
	partition := 0

	//连接至kafka集群的Leader节点
	conn, err := kafka.DialLeader(context.Background(), "tcp", "192.168.2.204:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	//设置发送消息的超时时间
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	//发送消息
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	//关闭连接
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func writeByWriter() {
	//创建一个writer，向topic发送消息
	w := &kafka.Writer{
		Addr:         kafka.TCP("192.168.2.204:9092"),
		Topic:        "kafka-test-topic",
		Balancer:     &kafka.LeastBytes{}, //指定分区的balancer模式为最小字节分布
		RequiredAcks: kafka.RequireAll,    //ack模式
		Async:        true,
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func writeByWriter3() {
	// 创建一个writer，向topic发送消息
	w := &kafka.Writer{
		Addr:         kafka.TCP("192.168.2.204:9092"),
		Topic:        "kafka-test-topic3",
		Balancer:     &kafka.LeastBytes{}, // 指定分区的balancer模式为最小字节分布
		RequiredAcks: kafka.RequireAll,    // ack模式
		Async:        true,
	}

	// 定义消息内容
	messages := []string{"Hello World!", "One!", "Two!", "song", "fa", "yuan"}

	// 循环发送消息
	for i, msg := range messages {
		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("Key-%d", i+1)), // 使用不同的分区键
				Value: []byte(msg),
			},
		)
		if err != nil {
			log.Fatalf("failed to write message: %v", err)
		}
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

// 创建不存在的topic
// 如果给Writer配置了AllowAutoTopicCreation:true，那么当发送消息至某个不存在的topic时，则会自动创建topic。
func writeByWriter2() {
	writer := kafka.Writer{
		Addr:                   kafka.TCP("192.168.2.204:9092"),
		Topic:                  "kafka-test-topic",
		AllowAutoTopicCreation: true, //自动创建topic
	}

	messages := []kafka.Message{
		{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		{
			Key:   []byte("Key-C"),
			Value: []byte("Tow!"),
		},
	}

	const retries = 3
	//重试3次
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := writer.WriteMessages(ctx, messages...)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			log.Fatal("unexpected error %v", err)
		}
		break
	}

	//关闭Writer
	if err := writer.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

// readByConn连接到kafka后接收消息
func readByConn() {
	//指定要连接的topic和partition
	topic := "kafka-test-topic"
	partition := 0

	//连接至kafka的Leader节点
	conn, err := kafka.DialLeader(context.Background(), "tcp", "192.168.2.204:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	//设置读取超时时间
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	//读取一批信息，得到的batch是一系列消息的迭代器
	batch := conn.ReadBatch(10e3, 1e6) //fetch 10KB min, 1MB max

	//遍历读取消息
	//b := make([]byte, 10e3) //10KB max per message
	fmt.Println("******遍历读取消息******")
	for {
		//使用batch.Read更高效一些，但是需要根据消息长度选择合适的buffer，如果传入的buffer太小（消息装不下）,就会返回io.ErrShortBuffer
		//n, err := batch.Read(b)
		//如果不考虑内存分配的效率问题，可以使用batch.ReadMessage读取消息
		mag, err := batch.ReadMessage()
		if err != nil {
			break
		}
		//fmt.Println(string(b[:n]))
		fmt.Println(string(mag.Value))
	}

	//关闭batch
	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	//关闭连接
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

// readByReader通过Reader接收消息
func readByReader() {
	//创建Reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"192.168.2.204:9092"},
		Topic:     "kafka-test-topic",
		Partition: 0,
		MaxBytes:  10e6, //10MB
	})
	//设置Offset
	reader.SetOffset(1)

	//接收消息
	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", message.Offset, string(message.Key), string(message.Value))
	}

	if err := reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

// 消费者组
func readByReaderGroup() {
	// 创建一个reader，指定GroupID，消费消息
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"192.168.2.204:9092"},
		GroupID:  "consumer-group-id", //指定消费者组ID
		Topic:    "kafka-test-topic",
		MaxBytes: 10e6, //10MB
	})

	//接收消息
	for {
		//ReadMessage会自动提交偏移量
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
	}

	//程序退出前关闭Reader
	if err := reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

// 消费者组，手动提交
func readByReaderGroup2() {
	// 创建一个reader，指定GroupID，消费消息
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"192.168.2.204:9092"},
		GroupID:  "consumer-group-id", //指定消费者组ID
		Topic:    "kafka-test-topic",
		MaxBytes: 10e6, //10MB
	})

	//接收消息
	ctx := context.Background()
	for {
		//获取消息
		message, err := reader.FetchMessage(ctx)
		if err != nil {
			break
		}
		//处理消息
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))

		//显示提交
		if err := reader.CommitMessages(ctx, message); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}

	//程序退出前关闭Reader
	if err := reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
