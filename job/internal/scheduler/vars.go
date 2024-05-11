package scheduler

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	MaxRetry = 3
)

// uniqueTaskId 唯一任务id
func uniqueTaskId(name string) asynq.Option {
	return asynq.TaskID(fmt.Sprintf("%s:%s", "job", name))
}

// 打印日志
func outputLog(funcName string, task any, entryID string, err error) {
	logPrefix := "task-scheduler"

	if err != nil {
		logx.Errorf("[task-scheduler][%s-%s] registered err : %+v , task : %+v", logPrefix, funcName, err, task)
	}

	if entryID != "" {
		logx.Debugf("[task-scheduler][%s-%s] entryID: %q", logPrefix, funcName, entryID)
	}
}

func sLog(funcName string, err error) {
	if err != nil {
		logx.Errorf("[任务调度][%s] %s", funcName, err.Error())
	}
}
