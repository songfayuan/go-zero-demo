package utils

import (
	"errors"
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"github.com/zeromicro/go-zero/rest/httpx"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"strings"
)

type RequestParam struct {
	Id int64 `json:"id"` // 用户id

	Name          string `json:"name,optional"`     // 账号
	NickName      string `json:"nickName,optional"` // 姓名
	UserName      string `json:"userName,optional"`
	JobName       string `json:"jobName,optional"`       // 职位名称
	AppName       string `json:"appName,optional"`       // 应用服务器名称
	AppSystemName string `json:"appSystemName,optional"` // 应用系统-名称
	RuleName      string `json:"ruleName,optional"`      // 规则名称
	PolicyName    string `json:"policyName,optional"`    // 策略名称
	LevelName     string `json:"levelName,optional"`
	SysName       string `json:"sysName,optional"`       // 归属系统-名称
	InventoryName string `json:"inventoryName,optional"` // 清单

	Email         string  `json:"email,optional"` // 邮箱
	Ids           []int64 `json:"ids,optional"`
	Password      string  `json:"password,optional"` // 新密码
	Status        int64   `json:"status,optional"`   // 用户状态  -1：禁用   1：正常
	Ip            string  `json:"ip,optional"`       // ip
	NewPassword   string  `json:"newPassword,optional"`
	RoleId        int64   `json:"roleId,optional"`
	MenuIds       []int64 `json:"menuIds,optional"`
	Label         string  `json:"label,optional"`  // 标签名
	Value         string  `json:"value,optional"`  // 数据值
	IdList        []int64 `json:"idList,optional"` // id集合
	Title         string  `json:"title,optional"`
	BackgroundUrl string  `json:"backgroundUrl,optional"` // 后台接口地址
	Url           string  `json:"url,optional"`           // 菜单路由
}

// GetIP returns request real ip.
func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}

func Post(request *http.Request) (*RequestParam, error) {
	var param RequestParam
	if err := httpx.ParseForm(request, &param); err != nil {
		return &RequestParam{}, err
	}
	if err := httpx.ParseJsonBody(request, &param); err != nil {
		return &RequestParam{}, err
	}
	return &param, nil
}

// 以1024作为基数
func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

// 将KB、MB、GB和TB转换为字节的
const (
	_        = iota
	KB int64 = 1 << (10 * iota)
	MB
	GB
	TB
)

// size是数值，uint是单位
func ToBytes(size int64, unit string) int64 {
	switch strings.ToLower(unit) {
	case "kb":
		return size * KB
	case "mb":
		return size * MB
	case "gb":
		return size * GB
	case "tb":
		return size * TB
	default:
		return size
	}
}

// GeoLite2IpCountry 根据ip获取国家名称
func GeoLite2IpCountry(ipAddr string) (country string) {
	// 使用 net 包中的函数判断是ip地址还是域名
	isIp := net.ParseIP(ipAddr)
	if isIp == nil { //非IP时，转换为ip
		ips, _ := net.LookupIP(ipAddr)
		ipAddr = ips[0].String()
	}

	// 指定相对路径的文件名
	filename := "dsms_admin/job/geolite2/GeoLite2-Country.mmdb"
	// 拼接相对路径和当前工作目录
	filePath := filepath.Join(filename)

	//db, err := geoip2.Open("iputil/geolite2/GeoLite2-ASN.mmdb")
	//db, err := geoip2.Open("iputil/geolite2/GeoLite2-City.mmdb")
	//db, err := geoip2.Open("common/utils/geolite2/GeoLite2-Country.mmdb")
	db, err := geoip2.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ip := net.ParseIP(ipAddr)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	country = record.Country.Names["en"]
	return country
}
