package kafkaconf

type Conf struct {
	Host        string
	Brokers     []string
	Group       string
	Topic       string
	Offset      string `json:",options=first|last,default=last"`
	OffsetId    int64  `json:",default=-1"` //-1时表示不使用该配置
	Consumers   int    `json:",default=8"`
	Processors  int    `json:",default=8"`
	MinBytes    int    `json:",default=10240"`    // 10K
	MaxBytes    int    `json:",default=10485760"` // 10M
	Username    string `json:",optional"`
	Password    string `json:",optional"`
	ForceCommit bool   `json:",default=true"`
}
