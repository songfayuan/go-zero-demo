package utils

import (
	"github.com/hibiken/asynq"
	goRedis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func NewAsynqClient(host, pass string) *asynq.Client {
	return asynq.NewClient(
		asynq.RedisClientOpt{Addr: host, Password: pass},
	)
}

func NewAsynqServer(host, pass string) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: host, Password: pass},
		asynq.Config{
			IsFailure: func(err error) bool {
				logx.Debugf("[task-job] job exec failure, \n\terr: %+v", err)
				return true
			},
			Concurrency: 30, //max concurrent process job task num
		},
	)
}

func NewScheduler(host, pass string) *asynq.Scheduler {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr:     host,
			Password: pass,
		}, &asynq.SchedulerOpts{
			Location: loc,
			PostEnqueueFunc: func(task *asynq.TaskInfo, err error) {
				logx.Debugf("[task-scheduler] scheduler failure, err: %+v, task: %+v\n", err, task)
			},
		})
}

func NewInspector(host, pass string) *asynq.Inspector {
	return asynq.NewInspector(
		asynq.RedisClientOpt{
			Addr:     host,
			Password: pass,
		})
}

// NewRedisNode
//
//	@Description:
//	@param host
//	@param pass
//	@param minIdleConn 空闲连接数
//	@return *goRedis.Client
func NewRedisNode(host, pass string, minIdleConn int) *goRedis.Client {
	return goRedis.NewClient(&goRedis.Options{
		Addr:         host,
		Password:     pass,
		MinIdleConns: minIdleConn,

		DialTimeout:           30 * time.Second,
		ReadTimeout:           30 * time.Second,
		WriteTimeout:          30 * time.Second,
		ContextTimeoutEnabled: false,

		MaxRetries: -1,

		PoolSize:        1000,
		PoolTimeout:     30 * time.Second,
		ConnMaxIdleTime: time.Minute,
	})
}
