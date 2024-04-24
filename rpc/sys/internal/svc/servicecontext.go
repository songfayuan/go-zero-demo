package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/common/task/kafkaconf"
	"go-zero-demo/rpc/model/sysmodel"
	"go-zero-demo/rpc/sys/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	KafkaConf *kafkaconf.Conf

	UserModel sysmodel.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	conf := redis.RedisConf{
		Host: c.RedisConf.Host,
		Type: c.RedisConf.Type,
		Pass: c.RedisConf.Pass,
		Tls:  c.RedisConf.Tls,
	}

	return &ServiceContext{
		Config:      c,
		RedisClient: redis.MustNewRedis(conf),

		KafkaConf: &c.KafkaConf,

		UserModel: sysmodel.NewSysUserModel(sqlConn),
	}
}
