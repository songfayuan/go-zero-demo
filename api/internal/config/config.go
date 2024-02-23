package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	SysRpc zrpc.RpcClientConf

	CacheRedis cache.ClusterConf

	Redis struct {
		Address string
		Pass    string
	}

	Mysql struct {
		Datasource string
	}
}
