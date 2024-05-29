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

type ClickhouseUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClickhouseUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClickhouseUpdateLogic {
	return &ClickhouseUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClickhouseUpdateLogic) ClickhouseUpdate(in *sysclient.ClickhouseReq) (*sysclient.ClickhouseResp, error) {
	// 更新数据
	updateData(l.ctx, l.svcCtx.ClickhouseConn, "Alice", "newpassword321")

	return &sysclient.ClickhouseResp{}, nil
}

// 更新数据
func updateData(ctx context.Context, conn clickhouse.Conn, name, newPassword string) {
	query := `
        ALTER TABLE demo UPDATE Password = ? WHERE Name = ?
    `
	err := conn.Exec(ctx, query, newPassword, name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data updated successfully")
}
