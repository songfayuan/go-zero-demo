package types

import "time"

type (
	// HttpModel http协议
	HttpModel struct {
		Id          string    `json:"id"`           //id
		KafkaOffset int64     `json:"kafka_offset"` //数据对应的kafka偏移id
		SnifferId   int64     `json:"sniffer_id"`   //探针id
		SnifferName string    `json:"sniffer_name"` //探针名称
		From        string    `json:"from"`         //
		PacketType  string    `json:"packet_type"`  //当前数据包的类型，如值为http表示就是http协议类型的数据包
		TotalLength int64     `json:"total_length"` //数据包流量大小，单位：字节
		CaptureTime time.Time `json:"capture_time"` //被捕获时的时间【Time转换而来】
		CreateTime  time.Time `json:"create_time"`  //数据添加时间

		Time    int64  `json:"time"`     //http请求数据包被捕获时的时间戳
		SrcAddr string `json:"src_addr"` //源mac地址
		DstAddr string `json:"dst_addr"` //目的mac地址
		Src     string `json:"src"`      //源IP和端口
		Dst     string `json:"dst"`      //目的IP和端口

		RequestMethod                 string `json:"request_method"`                   //http请求报文中请求行中的请求方法，如GET、POST等。
		RequestUrl                    string `json:"request_url"`                      //请求的统一资源标识符
		RequestHttpVersion            string `json:"request_http_version"`             //http协议版本
		RequestHost                   string `json:"request_host"`                     //接受请求的服务器地址，可以是IP或者是域名
		RequestContentType            string `json:"request_content-type"`             //请求内容类型
		RequestContentLength          string `json:"request_content-length"`           //请求内容长度
		RequestConnection             string `json:"request_connection"`               //指定与连接相关的属性，例如（Keep_Alive，长连接）
		RequestPragma                 string `json:"request_pragma"`                   //浏览器是否缓存资源
		RequestCacheControl           string `json:"request_cache-control"`            //请求缓存控制
		RequestAccept                 string `json:"request_accept"`                   //告诉服务器能够发送哪些媒体类型
		RequestUserAgent              string `json:"request_user-agent"`               //发送请求的应用名称
		RequestUpgrade                string `json:"request_upgrade"`                  //请求升级
		RequestOrigin                 string `json:"request_origin"`                   //请求的原始站点
		RequestSecWebsocketVersion    string `json:"request_sec-websocket-version"`    //请求的客户端所使用的的WebSocket协议版本号
		RequestSecWebsocketKey        string `json:"request_sec-websocket-key"`        //一个Base64编码值，由浏览器随机生成，用于升级request
		RequestSecWebsocketExtensions string `json:"request_sec-websocket-extensions"` //客户端想表达的协议级的扩展
		RequestAcceptEncoding         string `json:"request_accept-encoding"`          //通知服务器端可以发送的数据压缩格式
		RequestAcceptLanguage         string `json:"request_accept-language"`          //通知服务器端可以发送的语言
		RequestBody                   string `json:"request_body"`                     //请求体

		ResponseCode                   int64  `json:"response_code"`                     //状态码
		ResponseAcceptRanges           string `json:"response_accept-ranges"`            //服务器是否支持指定范围请求及哪种类型的分段请求
		ResponseAge                    string `json:"response_age"`                      //从原始服务器到代理缓存形成的估算时间
		ResponseCacheControl           string `json:"response_cache-control"`            //响应缓存控制
		ResponseReasonPhrase           string `json:"response_reason-phrase"`            //原因短语，数字状态码的可读版本，包含行终止序列之前的所有文本。原因短语只对人类有意义，因此，尽管响应行HTTP/1.0 200 NOT OK和HTTP/1.0 200 OK中原因短语的含义不同，但同样都会被当作成功指示处理
		ResponseServer                 string `json:"response_server"`                   //服务器应用软件的名称和版本
		ResponseXCache                 string `json:"response_x-cache"`                  //
		ResponseSecWebsocketAccept     string `json:"response_sec-websocket-accept"`     //服务器接受了客户端的请求
		ResponseSecWebsocketExtensions string `json:"response_sec-websocket-extensions"` //
		ResponseDate                   string `json:"response_date"`                     //提供日期和时间标志，说明回复报文是什么时间创建的
		ResponseLastModified           string `json:"response_last-modified"`            //响应最后一次修改时间
		ResponseContentType            string `json:"response_content-type"`             //响应正文的类型
		ResponseContentLength          string `json:"response_content-length"`           //响应正文的长度
		ResponseConnection             string `json:"response_connection"`               //允许客户端和服务器指定与请求/响应连接有关的选项
		ResponseTransferEncoding       string `json:"response_transfer-encoding"`        //告知接收端为了保证报文的可靠传输，对报文采用了什么编码方式
		ResponseBody                   string `json:"response_body"`                     //响应报文实体，该部分其实就是HTTP要传输的内容，是可选的

		ResponseXXssProtection string `json:"response_x-xss-protection"` //响应XSS保护机制
		ResponseXFrameOptions  string `json:"response_x-frame-options"`  //响应框架
	}

	// FtpModel ftp协议
	FtpModel struct {
		// 基本字段
		Id          int64     `json:"id"`           //id
		KafkaOffset int64     `json:"kafka_offset"` //数据对应的kafka偏移id
		SnifferId   int64     `json:"sniffer_id"`   //探针id
		SnifferName string    `json:"sniffer_name"` //探针名称
		From        string    `json:"from"`         //
		PacketType  string    `json:"packet_type"`  //当前数据包的类型，如值为http表示就是http协议类型的数据包
		TotalLength int64     `json:"total_length"` //数据包流量大小，单位：字节
		CaptureTime time.Time `json:"capture_time"` //被捕获时的时间【Time转换而来】
		CreateTime  time.Time `json:"create_time"`  //数据添加时间

		Time    int64  `json:"time"`     //http请求数据包被捕获时的时间戳
		SrcAddr string `json:"src_addr"` //源mac地址
		DstAddr string `json:"dst_addr"` //目的mac地址
		Src     string `json:"src"`      //源IP和端口
		Dst     string `json:"dst"`      //目的IP和端口

		// FtpModel ftp协议 字段
		FTime        string `json:"f_time"`        //格式化的时间
		SrcIp        string `json:"src_ip"`        //源ip
		DstIp        string `json:"dst_ip"`        //目的ip
		SrcPort      string `json:"src_port"`      //源端口
		DstPort      string `json:"dst_port"`      //目的端口
		Login        string `json:"login"`         //ftp服务器访问用户
		Pwd          string `json:"pwd"`           //访问用户密码
		TransferType string `json:"transfer_type"` //传送类型（上传，还是下载）"STOR"为上传文件，"RETR"为下载文件
		Filename     string `json:"filename"`      //传送的文件名
		Path         string `json:"path"`          //从ftp数据包中提取出的文件的保存路径
	}

	// MysqlOracleModel mysql、Oracle协议
	MysqlOracleModel struct {
		// 基本字段
		Id          int64     `json:"id"`           //id
		KafkaOffset int64     `json:"kafka_offset"` //数据对应的kafka偏移id
		SnifferId   int64     `json:"sniffer_id"`   //探针id
		SnifferName string    `json:"sniffer_name"` //探针名称
		From        string    `json:"from"`         //
		PacketType  string    `json:"packet_type"`  //当前数据包的类型，如值为http表示就是http协议类型的数据包
		TotalLength int64     `json:"total_length"` //数据包流量大小，单位：字节
		CaptureTime time.Time `json:"capture_time"` //被捕获时的时间【Time转换而来】
		CreateTime  time.Time `json:"create_time"`  //数据添加时间

		Time    int64  `json:"time"`     //http请求数据包被捕获时的时间戳
		SrcAddr string `json:"src_addr"` //源mac地址
		DstAddr string `json:"dst_addr"` //目的mac地址
		Src     string `json:"src"`      //源IP和端口
		Dst     string `json:"dst"`      //目的IP和端口

		// MysqlOracleModel mysql、Oracle协议
		SrcIp          string `json:"src_ip"`           //源ip
		DstIp          string `json:"dst_ip"`           //目的ip
		SrcPort        string `json:"src_port"`         //源端口
		DstPort        string `json:"dst_port"`         //目的端口
		User           string `json:"user"`             //数据库访问用户
		DBName         string `json:"db_name"`          //数据库名
		ReturnDataRows string `json:"return_data_rows"` //返回数据行数
		ConsumeTime    string `json:"consum_time"`      //耗时
		SqlExResult    string `json:"sql_ex_result"`    //sql执行结果状态："true"和"error"
		Sql            string `json:"sql"`              //sql执行语句
	}

	// SqlServerModel SqlServer协议
	SqlServerModel struct {
		// 基本字段
		Id          int64     `json:"id"`           //id
		KafkaOffset int64     `json:"kafka_offset"` //数据对应的kafka偏移id
		SnifferId   int64     `json:"sniffer_id"`   //探针id
		SnifferName string    `json:"sniffer_name"` //探针名称
		From        string    `json:"from"`         //
		PacketType  string    `json:"packet_type"`  //当前数据包的类型，如值为http表示就是http协议类型的数据包
		TotalLength int64     `json:"total_length"` //数据包流量大小，单位：字节
		CaptureTime time.Time `json:"capture_time"` //被捕获时的时间【Time转换而来】
		CreateTime  time.Time `json:"create_time"`  //数据添加时间

		Time    int64  `json:"time"`     //http请求数据包被捕获时的时间戳
		SrcAddr string `json:"src_addr"` //源mac地址
		DstAddr string `json:"dst_addr"` //目的mac地址
		Src     string `json:"src"`      //源IP和端口
		Dst     string `json:"dst"`      //目的IP和端口

		// SqlServerModel SqlServer协议
		SrcIp        string `json:"src_ip"`        //源ip
		DstIp        string `json:"dst_ip"`        //目的ip
		SrcPort      string `json:"src_port"`      //源端口
		DstPort      string `json:"dst_port"`      //目的端口
		StartTime    string `json:"start_time"`    //开始时间
		IpProtocol   int64  `json:"ip_protocol"`   //协议号
		ProtoStack   string `json:"proto_stack"`   //协议栈
		SessionId    int64  `json:"session_id"`    //会话ID
		SessionPkts  int64  `json:"session_pkts"`  //会话包数
		SessionBytes int64  `json:"session_bytes"` //会话字节数
		DBType       string `json:"DBType"`        //数据库类型
		DBIP         string `json:"DBIP"`          //数据库IP
		DBPort       string `json:"DBPort"`        //数据库端口
		PktType      string `json:"pktType"`       //包类型
		Log          string `json:"log"`           //数据库访问用户
		DBName       string `json:"DBName"`        //数据库名
		DBSQL        string `json:"DBSQL"`         //sql执行语句

		DbName         string `json:"db_name"`          //数据库名
		DbType         string `json:"dbtype"`           //数据库类型
		ReturnDataRows string `json:"return_data_rows"` //查询结果行数
		SqlExResult    string `json:"sql_ex_result"`    //查询结果状态
		Sql            string `json:"sql"`              //sql执行语句
	}

	// EmailModel email协议
	EmailModel struct {
		// 基本字段
		Id          int64     `json:"id"`           //id
		KafkaOffset int64     `json:"kafka_offset"` //数据对应的kafka偏移id
		SnifferId   int64     `json:"sniffer_id"`   //探针id
		SnifferName string    `json:"sniffer_name"` //探针名称
		From        string    `json:"from"`         //
		PacketType  string    `json:"packet_type"`  //当前数据包的类型，如值为http表示就是http协议类型的数据包
		TotalLength int64     `json:"total_length"` //数据包流量大小，单位：字节
		CaptureTime time.Time `json:"capture_time"` //被捕获时的时间【Time转换而来】
		CreateTime  time.Time `json:"create_time"`  //数据添加时间

		Time    int64  `json:"time"`     //http请求数据包被捕获时的时间戳
		SrcAddr string `json:"src_addr"` //源mac地址
		DstAddr string `json:"dst_addr"` //目的mac地址
		Src     string `json:"src"`      //源IP和端口
		Dst     string `json:"dst"`      //目的IP和端口

		// EmailModel email协议
		SrcIp        string `json:"src_ip"`        //源ip
		DstIp        string `json:"dst_ip"`        //目的ip
		SrcPort      int64  `json:"src_port"`      //源端口
		DstPort      int64  `json:"dst_port"`      //目的端口
		StartTime    string `json:"start_time"`    //开始时间
		IpProtocol   int64  `json:"ip_protocol"`   //协议号
		ProtoStack   string `json:"proto_stack"`   //协议栈
		SessionId    int64  `json:"session_id"`    //会话ID
		SessionPkts  int64  `json:"session_pkts"`  //会话包数
		SessionBytes int64  `json:"session_bytes"` //会话字节数
		SenderEmail  string `json:"senderEmail"`   //发件人邮箱
		SenderAli    string `json:"senderAli"`     //邮件发送者名字
		RcvrEmail    string `json:"rcvrEmail"`     //收件人邮箱
		RcvrAli      string `json:"rcvrAli"`       //邮件接收者名字
		Date         string `json:"date"`          //邮件时间字段中包含时间信息
		Subj         string `json:"subj"`          //邮件主题头
		MsgID        string `json:"msgID"`         //邮件MessageID头
		MsgIDCnt     int64  `json:"msgIDCnt"`      //邮件MessageID头计数
		Login        string `json:"login"`         //邮件登录的用户名
		Pwd          string `json:"pwd"`           //邮件登录的密码
		Content      string `json:"content"`       //邮件正文（纯文本）
		ContentLen   int64  `json:"contentLen"`    //邮件正文的内容长度
		Complemail   int64  `json:"complemail"`    //
		AttFileName  string `json:"attFileName"`   //邮件附件名称
	}

	// RedisModel Redis协议
	RedisModel struct {
		// 基本字段
		Id          int64     `json:"id"`           //id
		KafkaOffset int64     `json:"kafka_offset"` //数据对应的kafka偏移id
		SnifferId   int64     `json:"sniffer_id"`   //探针id
		SnifferName string    `json:"sniffer_name"` //探针名称
		From        string    `json:"from"`         //
		PacketType  string    `json:"packet_type"`  //当前数据包的类型，如值为http表示就是http协议类型的数据包
		TotalLength int64     `json:"total_length"` //数据包流量大小，单位：字节
		CaptureTime time.Time `json:"capture_time"` //被捕获时的时间【Time转换而来】
		CreateTime  time.Time `json:"create_time"`  //数据添加时间

		Time    int64  `json:"time"`     //http请求数据包被捕获时的时间戳
		SrcAddr string `json:"src_addr"` //源mac地址
		DstAddr string `json:"dst_addr"` //目的mac地址
		Src     string `json:"src"`      //源IP和端口
		Dst     string `json:"dst"`      //目的IP和端口

		// RedisModel Redis协议
		SrcIp          string `json:"src_ip"`          //源ip
		DstIp          string `json:"dst_ip"`          //目的ip
		SrcPort        int64  `json:"src_port"`        //源端口
		DstPort        int64  `json:"dst_port"`        //目的端口
		StartTime      string `json:"start_time"`      //开始时间
		RequestLength  int64  `json:"request_length"`  //请求数据长度,值的单位为字节
		ResponseLength int64  `json:"response_length"` //响应数据长度,值的单位为字节
		RequestType    string `json:"request_type"`    //请求数据类型
		ResponseType   string `json:"response_type"`   //响应数据类型
		Request        string `json:"request"`         //请求数据内容
		Response       string `json:"response"`        //响应数据内容
	}

	// DaMengModel 达梦数据库协议
	DaMengModel struct {
		// 基本字段
		Id          int64     `json:"id"`           //id
		KafkaOffset int64     `json:"kafka_offset"` //数据对应的kafka偏移id
		SnifferId   int64     `json:"sniffer_id"`   //探针id
		SnifferName string    `json:"sniffer_name"` //探针名称
		From        string    `json:"from"`         //
		PacketType  string    `json:"packet_type"`  //当前数据包的类型，如值为http表示就是http协议类型的数据包
		TotalLength int64     `json:"total_length"` //数据包流量大小，单位：字节
		CaptureTime time.Time `json:"capture_time"` //被捕获时的时间【Time转换而来】
		CreateTime  time.Time `json:"create_time"`  //数据添加时间

		Time    int64  `json:"time"`     //http请求数据包被捕获时的时间戳
		SrcAddr string `json:"src_addr"` //源mac地址
		DstAddr string `json:"dst_addr"` //目的mac地址
		Src     string `json:"src"`      //源IP和端口
		Dst     string `json:"dst"`      //目的IP和端口

		// DaMengModel 达梦数据库协议
		SrcIp       string `json:"src_ip"`        //源ip
		DstIp       string `json:"dst_ip"`        //目的ip
		SrcPort     string `json:"src_port"`      //源端口
		DstPort     string `json:"dst_port"`      //目的端口
		StartTime   string `json:"start_time"`    //开始时间
		User        string `json:"user"`          //数据库访问用户
		DBName      string `json:"db_name"`       //数据库名
		SqlExResult string `json:"sql_ex_result"` //sql执行结果状态
		Sql         string `json:"sql"`           //sql执行语句
	}

	// PostgresSqlModel postgresql 数据库协议
	PostgresSqlModel struct {
		// 基本字段
		Id          int64     `json:"id"`           //id
		KafkaOffset int64     `json:"kafka_offset"` //数据对应的kafka偏移id
		SnifferId   int64     `json:"sniffer_id"`   //探针id
		SnifferName string    `json:"sniffer_name"` //探针名称
		From        string    `json:"from"`         //
		PacketType  string    `json:"packet_type"`  //当前数据包的类型，如值为http表示就是http协议类型的数据包
		TotalLength int64     `json:"total_length"` //数据包流量大小，单位：字节
		CaptureTime time.Time `json:"capture_time"` //被捕获时的时间【Time转换而来】
		CreateTime  time.Time `json:"create_time"`  //数据添加时间

		Time    int64  `json:"time"`     //http请求数据包被捕获时的时间戳
		SrcAddr string `json:"src_addr"` //源mac地址
		DstAddr string `json:"dst_addr"` //目的mac地址
		Src     string `json:"src"`      //源IP和端口
		Dst     string `json:"dst"`      //目的IP和端口

		// PostgresSqlModel postgresql 数据库协议
		SrcIp       string `json:"src_ip"`        //源ip
		DstIp       string `json:"dst_ip"`        //目的ip
		SrcPort     int64  `json:"src_port"`      //源端口
		DstPort     int64  `json:"dst_port"`      //目的端口
		StartTime   string `json:"start_time"`    //开始时间
		DbType      string `json:"dbtype"`        //数据库类型
		User        string `json:"user"`          //数据库访问用户
		DBName      string `json:"db_name"`       //数据库名
		SqlExResult string `json:"sql_ex_result"` //sql执行结果状态
		Sql         string `json:"sql"`           //sql执行语句
	}
)
