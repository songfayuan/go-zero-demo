package scheduler

import (
	"github.com/hibiken/asynq"
	jobtype "go-zero-demo/job/internal/const"
)

// scheduler job ------> ./app/task/job/internal/logic/demoJob.go.
func (l *TaskScheduler) demoScheduler() {
	jobName := jobtype.TestDemoJob
	if !l.canRun(jobName) {
		return
	}

	task := asynq.NewTask(jobName, nil)

	// 定时循环
	//entryID, err := l.svcCtx.Scheduler.Register("* * * * * *", task, asynq.Unique(120*time.Second))
	//entryID, err := l.svcCtx.Scheduler.Register("@every 1h30m", task, asynq.Unique(50*time.Second), asynq.MaxRetry(3))
	//_, err := l.svcCtx.AsynqScheduler.Register("@every 10s", task)
	//每1分钟执行一次任务
	_, err := l.svcCtx.AsynqScheduler.Register("*/1 * * * *", task, uniqueTaskId(jobName), asynq.MaxRetry(MaxRetry))

	// 一次性任务
	//_, err := l.svcCtx.AsynqClient.Enqueue(task, uniqueTaskId(jobName), asynq.MaxRetry(MaxRetry))
	// 一次性任务 延时执行
	//taskInfo, err = l.svcCtx.AsynqClient.Enqueue(task, uniqueTaskId(jobName), asynq.ProcessIn(3*time.Second))
	// 一次性任务 指定时间
	//taskInfo, err = l.svcCtx.AsynqClient.Enqueue(task, uniqueTaskId(jobName), asynq.ProcessAt(time.Now().Add(time.Second)))

	sLog("demoScheduler", err)
}
