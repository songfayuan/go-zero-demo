package scheduler

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/job/internal/config"
	"go-zero-demo/job/internal/svc"
	"time"
)

type TaskScheduler struct {
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	jobConf *config.JobConf
}

// Register 注册业务调度
func (l *TaskScheduler) Register() {
	logx.Info("注册业务调度....")
	l.runAllArchivedTasks() // 检测并运行异常任务

	// 示例
	l.demoScheduler()
}

func (l *TaskScheduler) runAllArchivedTasks() {
	ee := l.svcCtx.AsynqInspector.DeleteQueue("default", true)
	if ee != nil {
		logx.Errorf("[调度任务] DeleteQueue err: %s", ee.Error())
	}

	go func() {
		ticker := time.NewTicker(15 * time.Second)
		for range ticker.C {
			// 运行上次异常的
			list, err := l.svcCtx.AsynqInspector.ListArchivedTasks("default")
			if err != nil {
				logx.Errorf("[调度任务] 获取重试列表错误 err: %s", err.Error())
			}

			if len(list) > 0 {
				logx.Errorf("[调度任务-重试] 重试任务数：%d", len(list))
				for _, t := range list {
					e := l.svcCtx.AsynqInspector.RunTask(t.Queue, t.ID)
					if e != nil {
						logx.Infof("[调度任务-重试] %s-执行失败: %s", t.Type, e.Error())
					} else {
						logx.Infof("[调度任务-重试] %s-执行成功", t.Type)
					}
				}
			}
		}
	}()
}

func NewCronScheduler(ctx context.Context, svcCtx *svc.ServiceContext, jobConf *config.JobConf) *TaskScheduler {
	return &TaskScheduler{
		ctx:     ctx,
		svcCtx:  svcCtx,
		jobConf: jobConf.New(),
	}
}

// 检测是否要运行
func (l *TaskScheduler) canRun(jobName string) bool {
	if l.jobConf.Disable != nil && len(l.jobConf.Disable) > 0 {
		// 判断禁用列表
		for _, i := range l.jobConf.Disable {
			if jobName == i {
				logx.Infof("[调度任务-stop] 在禁用名单中: %s", jobName)
				return false
			}
		}

		logx.Infof("[调度任务-run] 不在禁用名单中：%s", jobName)
		return true

	} else if l.jobConf.Enable != nil && len(l.jobConf.Enable) > 0 {
		// 判断启用列表
		for _, i := range l.jobConf.Enable {
			if jobName == i {
				logx.Infof("[调度任务-run] 在运行名单中：%s", jobName)
				return true
			}
		}

		logx.Infof("[调度任务-stop] 不在运行名单中：%s", jobName)
		return false

	}

	logx.Infof("[调度任务-run] 未配置-放行：%s", jobName)
	return true
}
