package logic

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"log"

	"go-zero-demo/rpc/sys/internal/svc"
	"go-zero-demo/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClickhouseDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClickhouseDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClickhouseDeleteLogic {
	return &ClickhouseDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClickhouseDeleteLogic) ClickhouseDelete(in *sysclient.ClickhouseReq) (*sysclient.ClickhouseResp, error) {
	// 删除数据
	deleteData(l.ctx, l.svcCtx.ClickhouseConn, "Bob")

	return &sysclient.ClickhouseResp{}, nil
}

// 删除数据
func deleteData(ctx context.Context, conn clickhouse.Conn, name string) {
	query := `
        ALTER TABLE demo DELETE WHERE Name = ?
    `
	err := conn.Exec(ctx, query, name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("删除数据成功...")
}
