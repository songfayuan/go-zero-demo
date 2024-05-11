package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	jobtype "go-zero-demo/job/internal/const"
	"go-zero-demo/job/internal/svc"
	"runtime/debug"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register 注册业务逻辑
func (l *CronJob) Register() *asynq.ServeMux {
	logx.Info("注册业务逻辑....")
	mux := asynq.NewServeMux()

	// demo job
	mux.Handle(jobtype.TestDemoJob, NewDemoJobHandler(l.svcCtx))

	return mux
}

// RunSimpleJob 直接调用的程序
func RunSimpleJob(ctx context.Context, svcContext *svc.ServiceContext) {
	defer func() {
		if r := recover(); r != nil {
			s := string(debug.Stack())
			logx.Errorf("[RunSimpleJob] %v, stack=%s", r, s)
		}
	}()
	// todo 横向扩展服务时应根据节点类型判断是否需要运行以下服务
	logx.Info("这里是直接运行的程序....")

}
