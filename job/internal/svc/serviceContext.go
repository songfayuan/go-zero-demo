package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-demo/common/utils"
	"go-zero-demo/job/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	/* asynq 任务组件 */
	AsynqScheduler *asynq.Scheduler //调度器
	AsynqInspector *asynq.Inspector //检查器客户端
	AsynqServer    *asynq.Server    //服务端
	AsynqClient    *asynq.Client    //客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
	conf := redis.RedisConf{
		Host: c.RedisConf.Host,
		Type: c.RedisConf.Type,
		Pass: c.RedisConf.Pass,
		Tls:  c.RedisConf.Tls,
	}
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.MustNewRedis(conf),

		AsynqScheduler: utils.NewScheduler(conf.Host, conf.Pass),
		AsynqInspector: utils.NewInspector(conf.Host, conf.Pass),
		AsynqServer:    utils.NewAsynqServer(conf.Host, conf.Pass),
		AsynqClient:    utils.NewAsynqClient(conf.Host, conf.Pass),
	}
}
