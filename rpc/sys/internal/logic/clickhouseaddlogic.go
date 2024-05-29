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

type ClickhouseAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type User struct {
	Name     string
	NickName string
	Password string
	Email    string
}

func NewClickhouseAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClickhouseAddLogic {
	return &ClickhouseAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// clickhouse增删改查
func (l *ClickhouseAddLogic) ClickhouseAdd(in *sysclient.ClickhouseReq) (*sysclient.ClickhouseResp, error) {
	// 创建表
	createTable(l.ctx, l.svcCtx.ClickhouseConn)

	// 插入数据
	insertData(l.ctx, l.svcCtx.ClickhouseConn, "Alice", "Ally", "password123", "alice@example.com")
	insertData(l.ctx, l.svcCtx.ClickhouseConn, "Bob", "Bobby", "password456", "bob@example.com")

	// 批量插入数据
	batchInsertData(l.ctx, l.svcCtx.ClickhouseConn, []User{
		{"Alice", "Ally", "password123", "alice@example.com"},
		{"Bob", "Bobby", "password456", "bob@example.com"},
		{"Charlie", "Char", "password789", "charlie@example.com"},
	})

	return &sysclient.ClickhouseResp{}, nil
}

// 创建表
func createTable(ctx context.Context, conn clickhouse.Conn) {
	query := `
        CREATE TABLE IF NOT EXISTS demo (
            Name String,
            NickName String,
            Password String,
            Email String
        ) ENGINE = MergeTree()
        ORDER BY Name
    `
	err := conn.Exec(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功创建demo表...")
}

// 插入数据
func insertData(ctx context.Context, conn clickhouse.Conn, name, nickName, password, email string) {
	batch, err := conn.PrepareBatch(ctx, "INSERT INTO demo (Name, NickName, Password, Email)")
	if err != nil {
		log.Fatal(err)
	}

	err = batch.Append(name, nickName, password, email)
	if err != nil {
		log.Fatal(err)
	}

	err = batch.Send()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功插入数据...")
}

// 批量插入数据
func batchInsertData(ctx context.Context, conn clickhouse.Conn, users []User) {
	batch, err := conn.PrepareBatch(ctx, "INSERT INTO demo (Name, NickName, Password, Email)")
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		err := batch.Append(user.Name, user.NickName, user.Password, user.Email)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = batch.Send()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功批量插入数据...")
}
