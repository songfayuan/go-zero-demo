package utils

import (
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

const (
	//溯源任务表名
	TrackSourceTaskTableName string = "dsms_track_source_task"
	//文档指纹表
	DocFingerPrintTableName string = "dsms_flow_doc_fingerprint"
	//文档指纹子表
	DocFingerPrintDocTableName string = "dsms_flow_doc_fingerprint_doc"
	//系统管理—证书管理表
	CertificateTableName string = "dsms_sys_certificate"
	//系统管理—自定义管理表
	CustomizeTableName string = "dsms_sys_customize"
)

// 判断文件夹是否存在
func HasDir(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 创建文件夹
func CreateDir(path string) {
	exist, err := HasDir(path)
	if err != nil {
		logx.Infof("获取文件夹异常: %+v", err)
		return
	}
	if !exist {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			logx.Infof("创建文件夹异常!: %+v", err)
		} else {
			logx.Infof("创建文件夹成功!")
		}
	} else {
		//logx.Infof("文件夹已存在! %+v", err)
	}
}

// Jaccard相似系数：比较内容相似度
func JaccardSimilarity(s1, s2 string) float64 {
	set1 := make(map[rune]bool)
	set2 := make(map[rune]bool)
	for _, r := range s1 {
		set1[r] = true
	}
	for _, r := range s2 {
		set2[r] = true
	}
	intersection := 0
	for r := range set1 {
		if set2[r] {
			intersection++
		}
	}
	union := len(set1) + len(set2) - intersection
	return float64(intersection) / float64(union)
}
