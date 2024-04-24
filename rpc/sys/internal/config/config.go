package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/common/task/kafkaconf"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		Datasource string
	}

	RedisConf struct {
		Host string
		Type string `json:",default=node,options=node|cluster"`
		Pass string `json:",optional"`
		Tls  bool   `json:",optional"`
	}

	KafkaConf kafkaconf.Conf
}
