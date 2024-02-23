package utils

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	goRedis "github.com/redis/go-redis/v9"
	"time"
)

func Max[T int | int64](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Min[T int | int64](a, b T) T {
	if a > b {
		return b
	}

	return a
}

func DelayUnlock(mutex *redsync.Mutex, client *goRedis.Client) {
	name, val := mutex.Name(), mutex.Value()

	// 解锁
	_, _ = mutex.Unlock()

	// 再加个锁
	client.SetNX(context.Background(), name, val, 1*time.Second)
}
