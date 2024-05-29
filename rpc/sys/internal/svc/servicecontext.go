package svc

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/common/task/kafkaconf"
	"go-zero-demo/rpc/model/sysmodel"
	"go-zero-demo/rpc/sys/internal/config"
	"log"
	"time"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	KafkaConf      *kafkaconf.Conf
	ClickhouseConn clickhouse.Conn

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

		KafkaConf:      &c.KafkaConf,
		ClickhouseConn: NewClickHouseConn(c.ClickHouse),

		UserModel: sysmodel.NewSysUserModel(sqlConn),
	}
}

// NewClickHouseConn 创建clickhouse链接 【clickhouse-go】
func NewClickHouseConn(ckConfig config.ClickHouseConf) clickhouse.Conn {
	maxIdleConns := 20
	if ckConfig.Opt2.MaxIdleConns > 0 {
		maxIdleConns = ckConfig.Opt2.MaxIdleConns
	}
	maxOpenConns := 2000
	if ckConfig.Opt2.MaxOpenConns > 0 {
		maxOpenConns = ckConfig.Opt2.MaxOpenConns
	}

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: ckConfig.Addr,
		Auth: clickhouse.Auth{
			Database: ckConfig.Auth.Database,
			Username: ckConfig.Auth.Username,
			Password: ckConfig.Auth.Password,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout:  30 * time.Second,
		MaxIdleConns: maxIdleConns,
		MaxOpenConns: maxOpenConns,
		Debug:        ckConfig.Debug,
	})

	if err != nil {
		log.Fatalf("error: 启动clickhouse client失败, %v", err)
	}

	return conn
}
