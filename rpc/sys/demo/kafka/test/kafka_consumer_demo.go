package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
	"time"
)

// 多个消费者同时消费同一Topic的数据
// 一个partition，只能被消费组里的一个消费者消费，但是可以同时被多个消费组消费
func main() {
	// 创建消费者组ID
	consumerGroupID := "consumer-group-id16"

	// 创建两个消费者
	consumer1 := createConsumer(consumerGroupID, "Consumer1")
	consumer2 := createConsumer(consumerGroupID, "Consumer2")
	consumer3 := createConsumer(consumerGroupID, "Consumer3")
	consumer4 := createConsumer(consumerGroupID, "consumer4")
	consumer5 := createConsumer(consumerGroupID, "consumer5")

	// 启动消费者
	var wg sync.WaitGroup
	wg.Add(4)
	go consumeMessages(consumer1, &wg, "Consumer1")
	go consumeMessages(consumer2, &wg, "Consumer2")
	go consumeMessages(consumer3, &wg, "consumer3")
	go consumeMessages(consumer4, &wg, "consumer4")
	go consumeMessages(consumer5, &wg, "consumer5")

	wg.Wait()
}

func createConsumer(groupID, consumerName string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"192.168.2.204:9092"},
		GroupID: groupID,
		Topic:   "kafka-test-topic3",
	})
}

func consumeMessages(reader *kafka.Reader, wg *sync.WaitGroup, consumerName string) {
	defer wg.Done()
	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		time.Sleep(1 * time.Second)

		fmt.Printf("[%s] Message at topic/partition/offset %v/%v/%v: %s = %s\n", consumerName, message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
	}
	if err := reader.Close(); err != nil {
		log.Fatalf("[%s] Failed to close reader: %v\n", consumerName, err)
	}
}
