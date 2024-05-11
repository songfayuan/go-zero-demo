package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/job/internal/config"
	"go-zero-demo/job/internal/logic"
	"go-zero-demo/job/internal/scheduler"
	"go-zero-demo/job/internal/svc"
	"time"
)

var configFile = flag.String("f", "etc/job.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	sqlx.DisableStmtLog() // 禁用记录sql日志
	logx.DisableStat()    // 禁用记录资源日志

	c.MustSetUp()

	ctx := context.Background()
	svcContext := svc.NewServiceContext(c)

	// 直接运行的程序
	logic.RunSimpleJob(ctx, svcContext)

	waitTime := time.Second * 5
	if c.Mode == service.DevMode {
		waitTime = time.Millisecond * 300
	}
	time.AfterFunc(waitTime, func() {
		taskScheduler(ctx, svcContext, c)
	})

	taskJob(ctx, svcContext)
}

// 启动多个并发工作线程来处理客户端创建的任务
func taskJob(ctx context.Context, svcContext *svc.ServiceContext) {
	cronJob := logic.NewCronJob(ctx, svcContext)
	mux := cronJob.Register()

	fmt.Printf("[task-job] starting server\n\n")

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("[task-job] run err : %+v", err)
	}
}

// 创建并安排由后台工作人员异步处理的任务
func taskScheduler(ctx context.Context, svcContext *svc.ServiceContext, c config.Config) {
	cronScheduler := scheduler.NewCronScheduler(ctx, svcContext, c.JobConf.New())
	cronScheduler.Register()

	fmt.Printf("[task-scheduler] starting server\n\n")

	if err := svcContext.AsynqScheduler.Run(); err != nil {
		logx.WithContext(ctx).Errorf("[task-scheduler] run err : %+v", err)
	}
}
