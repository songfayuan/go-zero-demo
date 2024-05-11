package logic

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"go-zero-demo/job/internal/svc"
	"time"
)

type DemoJobHandler struct {
	svcCtx *svc.ServiceContext
}

func NewDemoJobHandler(svcCtx *svc.ServiceContext) *DemoJobHandler {
	return &DemoJobHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask if return err != nil , asynq will retry
func (l *DemoJobHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {

	fmt.Printf("demo job -----> %s\n", time.Now())

	return nil
}
