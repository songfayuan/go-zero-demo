package utils

import (
	"fmt"
	"sync"
	"time"
)

const (
	// 机器ID位数
	workerBits = 5
	// 数据中心ID位数
	datacenterBits = 5
	// 序列号位数
	sequenceBits = 12

	// 最大机器ID和数据中心ID
	maxWorkerID     = -1 ^ (-1 << workerBits)
	maxDatacenterID = -1 ^ (-1 << datacenterBits)
	// 最大序列号
	maxSequenceMask = -1 ^ (-1 << sequenceBits)

	// 机器ID向左移位数
	workerShift = sequenceBits
	// 数据中心ID向左移位数
	datacenterShift = workerBits + workerShift
	// 时间戳向左移位数
	timestampShift = datacenterBits + datacenterShift

	// 自定义起始时间（这个值可以根据需要调整）
	customStartTime = int64(1136185445000) // 2006-01-02 15:04:05
)

// Snowflake 结构
type Snowflake struct {
	mu            sync.Mutex
	timestamp     int64
	workerID      int64
	datacenterID  int64
	sequence      int64
	lastTimestamp int64
}

// NewSnowflake 创建一个Snowflake实例
func NewSnowflake(workerID, datacenterID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, fmt.Errorf("Worker ID must be between 0 and %d", maxWorkerID)
	}

	if datacenterID < 0 || datacenterID > maxDatacenterID {
		return nil, fmt.Errorf("Datacenter ID must be between 0 and %d", maxDatacenterID)
	}

	return &Snowflake{
		timestamp:     0,
		workerID:      workerID,
		datacenterID:  datacenterID,
		sequence:      0,
		lastTimestamp: -1,
	}, nil
}

// NextID 生成下一个唯一ID
func (sf *Snowflake) NextID() (int64, error) {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	currentTimestamp := sf.timeGen()

	if currentTimestamp < sf.lastTimestamp {
		return 0, fmt.Errorf("Clock moved backwards. Refusing to generate ID")
	}

	if currentTimestamp == sf.lastTimestamp {
		sf.sequence = (sf.sequence + 1) & maxSequenceMask
		if sf.sequence == 0 {
			// 序列号溢出，等待下一个时间戳
			currentTimestamp = sf.waitNextMillis(currentTimestamp)
		}
	} else {
		sf.sequence = 0
	}

	sf.lastTimestamp = currentTimestamp

	// 使用位运算生成最终的Snowflake ID
	ID := ((currentTimestamp - customStartTime) << timestampShift) |
		(sf.datacenterID << datacenterShift) |
		(sf.workerID << workerShift) |
		(sf.sequence)

	return ID, nil
}

// timeGen 获取当前时间戳
func (sf *Snowflake) timeGen() int64 {
	return time.Now().UnixNano() / 1000000
}

// waitNextMillis 自旋等待下一个时间戳
func (sf *Snowflake) waitNextMillis(lastTimestamp int64) int64 {
	currentTimestamp := sf.timeGen()
	for currentTimestamp <= lastTimestamp {
		currentTimestamp = sf.timeGen()
	}
	return currentTimestamp
}

func main() {
	// 创建Snowflake实例，传入机器ID和数据中心ID
	sf, err := NewSnowflake(1, 1)
	if err != nil {
		fmt.Printf("Error creating Snowflake: %s\n", err)
		return
	}

	// 生成并打印一些唯一ID
	for i := 0; i < 5; i++ {
		id, err := sf.NextID()
		if err != nil {
			fmt.Printf("Error generating ID: %s\n", err)
		} else {
			fmt.Printf("Generated ID: %d\n", id)
		}
	}
}
