package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"strings"
)

type Config struct {
	service.ServiceConf

	RedisConf struct {
		Host string
		Type string `json:",default=node,options=node|cluster"`
		Pass string `json:",optional"`
		Tls  bool   `json:",optional"`
	}

	JobConf JobConf

	TaskSet struct {
		StatFlowRecordLogic struct {
			HttpCaptureFilterCheckSecond int64  `json:",optional"`
			JobFlowGroupId               string `json:",optional"`
			CommitMsgBatchCount          int64  `json:",optional"`
			CommitMsgBatchSecond         int64  `json:",optional"`
			CommitIPBatchSecond          int64  `json:",optional"`
		} `json:",optional"`
	}
}

// JobConf 优先判断是否包含禁用列表Disable，若已配置Disable，则不判断Enable
type JobConf struct {
	Disable []string `json:",optional"` // 禁用的任务
	Enable  []string `json:",optional"` // 启用的任务（Disable为空时生效）
}

func (j *JobConf) New() *JobConf {
	if j.Disable != nil {
		disable := make([]string, 0)
		for _, d := range j.Disable {
			if newD := strings.Trim(d, " "); newD != "" {
				disable = append(disable, newD)
			}
		}
		j.Disable = disable
	}

	if (j.Disable == nil || len(j.Disable) == 0) && j.Enable != nil {
		able := make([]string, 0)
		for _, e := range j.Enable {
			if newE := strings.Trim(e, " "); newE != "" {
				able = append(able, newE)
			}
		}
		j.Enable = able
	} else {
		j.Enable = nil
	}
	return j
}
