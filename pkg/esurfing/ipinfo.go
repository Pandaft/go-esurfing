package esurfing

import (
	"errors"
	"github.com/Pandaft/go-esurfing/internal/logger"
	"github.com/imroc/req/v3"
	"regexp"
	"time"
)

type ESurfIpInfo struct {
	NasIp    string `json:"nasip"`
	ClientIp string `json:"clientip"`
}

const (
	getESurfIpInfoUrl         = "http://189.cn/"
	getESurfIpInfoUrlRedirect = "https://www.189.cn/"
)

var (
	nasIpRegex    = regexp.MustCompile(`wlanacip=([\d\.]+)&?`)
	clientIpRegex = regexp.MustCompile(`wlanuserip=([\d\.]+)&?`)
)

// GetESurfIpInfo 获取 IP 信息
func GetESurfIpInfo() (esurfIpInfo ESurfIpInfo, err error) {

	log := logger.GetLogger("获取IP信息")

	// 记录耗时
	start := time.Now()
	log.Debug("开始")
	defer func() {
		elapsed := time.Now().Sub(start)
		log.Debugf("结束，耗时：%s", elapsed)
	}()

	// 新客户端
	client := req.C()
	client.SetRedirectPolicy(req.NoRedirectPolicy())

	// 发送请求
	log.Debugf("发送请求")
	resp, err := client.R().Get(getESurfIpInfoUrl)

	// 发生错误
	if err != nil {
		log.Debugf("发生错误：%s", err)
		return
	}

	// 获取重定向链接
	location, err := resp.Location()
	if err != nil {
		log.Debugf("获取重定向链接错误：%s", err)
		return
	}
	redirectUrl := location.String()
	log.Debugf("重定向链接：%s", redirectUrl)

	// 检查是否已经登录（通过是否已联网判断）
	if redirectUrl == getESurfIpInfoUrlRedirect {
		err = errors.New("当前已登入")
		log.Debug(err)
		return
	}

	// 提取 IP 信息
	match := nasIpRegex.FindStringSubmatch(redirectUrl)
	if len(match) < 2 {
		err = errors.New("未能从重定向链接中提取 nasip 参数")
		return
	}
	esurfIpInfo.NasIp = match[1]
	log.Debugf("提取 nasip 参数成功（%s）", match[1])
	match = clientIpRegex.FindStringSubmatch(redirectUrl)
	if len(match) < 2 {
		err = errors.New("未能从重定向链接中提取 clientip 参数")
		return
	}
	esurfIpInfo.ClientIp = match[1]
	log.Debugf("提取 clientip 参数成功（%s）", match[1])

	return
}
