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

	ClickHouse ClickHouseConf
}

type ClickHouseAuthConf struct {
	Database string
	Username string
	Password string `json:",optional"`
}

type ClickHouseConf struct {
	Addr []string
	Auth ClickHouseAuthConf
	Opt1 struct {
		MaxIdleConns int `json:",optional"`
		MaxOpenConns int `json:",optional"`
	} `json:",optional"`
	Opt2 struct {
		MaxIdleConns int `json:",optional"`
		MaxOpenConns int `json:",optional"`
	} `json:",optional"`
	Table   string
	Columns []string `json:",optional"`
	Debug   bool     `json:",default=false"`

	Datasource string `json:",optional,default="`
}
