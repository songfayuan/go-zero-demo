package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/rpc/model/sysmodel"
	"go-zero-demo/rpc/sys/internal/config"
)

type ServiceContext struct {
	Config config.Config

	UserModel sysmodel.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	return &ServiceContext{
		Config: c,

		UserModel: sysmodel.NewSysUserModel(sqlConn),
	}
}
