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

type ClickhouseGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClickhouseGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClickhouseGetLogic {
	return &ClickhouseGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClickhouseGetLogic) ClickhouseGet(in *sysclient.ClickhouseReq) (*sysclient.ClickhouseResp, error) {
	// 查询数据
	queryData(l.ctx, l.svcCtx.ClickhouseConn)

	return &sysclient.ClickhouseResp{}, nil
}

// 查询数据
func queryData(ctx context.Context, conn clickhouse.Conn) {
	rows, err := conn.Query(ctx, "SELECT Name, NickName, Password, Email FROM demo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, nickName, password, email string
		if err := rows.Scan(&name, &nickName, &password, &email); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Name: %s, NickName: %s, Password: %s, Email: %s\n", name, nickName, password, email)
	}
}
