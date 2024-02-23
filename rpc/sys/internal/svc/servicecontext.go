package svc

import (
	"errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"
	"go-zero-demo/rpc/model/sysmodel"
	"go-zero-demo/rpc/sys/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
	Redis  *redis.Redis

	UserModel sysmodel.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	ca := cache.New(c.CacheRedis, syncx.NewSingleFlight(), cache.NewStat("dc"), errors.New("data not find"))
	rConn := redis.New(c.CacheRedis[0].Host, func(r *redis.Redis) {
		r.Type = c.CacheRedis[0].Type
		r.Pass = c.CacheRedis[0].Pass
	})

	return &ServiceContext{
		Config: c,
		Cache:  ca,
		Redis:  rConn,

		UserModel: sysmodel.NewSysUserModel(sqlConn),
	}
}
