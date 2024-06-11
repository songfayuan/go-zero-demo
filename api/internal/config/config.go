package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	SysRpc zrpc.RpcClientConf

	Mysql struct {
		Datasource string
	}

	UploadFile UploadFile
}

type UploadFile struct {
	MaxFileNum   int64
	MaxFileSize  int64
	SavePath     string
	TemplatePath string
}
