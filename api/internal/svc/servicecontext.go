package svc

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/api/internal/config"
	"go-zero-demo/rpc/sys/sys"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ServiceContext struct {
	Config config.Config

	Sys   sys.Sys
	Redis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	return &ServiceContext{
		Config: c,
		Sys:    sys.NewSys(zrpc.MustNewClient(c.SysRpc, zrpc.WithUnaryClientInterceptor(interceptor))),
		Redis:  newRedis,
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}

func interceptor(ctx context.Context, method string, req any, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md := metadata.New(map[string]string{"x": "xx"})
	ctx = metadata.NewOutgoingContext(ctx, md)
	// logx.Debug("调用rpc服务前")
	err := invoker(ctx, method, req, reply, cc)
	if err != nil {
		return err
	}
	// logx.Debug("调用rpc服务后")
	return nil
}
