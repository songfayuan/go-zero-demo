package types

type IndexSeriesData struct {
	Name  string   `json:"name"`  // 名称
	YData []int64  `json:"yData"` // Y轴数据-访问次数
	XData []string `json:"xData"` // X轴数据-访问时间
}

type Pie struct {
	Name  string `json:"name"`  // 名称
	Value int64  `json:"value"` // 数据
}

type IndexPieData struct {
	Name string `json:"name"` // 名称
	List []*Pie `json:"list"` // 数据集合
}

type Analysis struct {
	AppName            string `json:"appName"`            // 应用名称
	AllTerminal        int64  `json:"allTerminal"`        // 所有终端数量
	AccessTerminal     int64  `json:"accessTerminal"`     // 访问过应用的终端数量
	AllAccount         int64  `json:"allAccount"`         // 所有账号数量
	AccessAccount      int64  `json:"accessAccount"`      // 访问过应用的账号数量
	AllLinkApi         int64  `json:"allLinkApi"`         // 所有接口数量
	AccessLinkApi      int64  `json:"accessLinkApi"`      // 访问过应用的接口数量
	AllSensitiveApi    int64  `json:"allSensitiveApi"`    // 所有敏感接口数量
	AccessSensitiveApi int64  `json:"accessSensitiveApi"` // 调用过的敏感接口数量
	Request            int64  `json:"request"`            // 请求数量
	Response           string `json:"response"`           // 响应流量
}

type OutFlow struct {
	Index    int64  `json:"index"`    // 顺序
	YData    string `json:"yData"`    // Y轴数据-IP地址信息
	XData    int64  `json:"xData"`    // X轴数据-流量大小
	ShowData string `json:"showData"` // X轴数据-显示的数据
}

type DisplayData struct {
	FlowTrendList     []*IndexSeriesData `json:"flowTrendList"`     // 流量趋势
	RiskTrendList     []*IndexSeriesData `json:"riskTrendList"`     // 风险趋势
	LevelStructured   *IndexPieData      `json:"levelStructured"`   // 分级-结构化
	LevelUnStructured *IndexPieData      `json:"levelUnStructured"` // 分级-非结构化
	ClassStructured   *IndexPieData      `json:"classStructured"`   // 分类-结构化
	ClassUnStructured *IndexPieData      `json:"classUnStructured"` // 分类-非结构化
	AnalysisData      []*Analysis        `json:"analysisData"`      // 应用分析
	OutFlowData       []*OutFlow         `json:"outFlowData"`       // 流出流量TOP
}
