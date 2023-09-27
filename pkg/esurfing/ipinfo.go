package esurfing

import (
	"errors"
	"github.com/charmbracelet/log"
	"github.com/imroc/req/v3"
	"regexp"
)

type ESurfIpInfo struct {
	NasIp    string
	ClientIp string
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

	client := req.C()
	client.SetRedirectPolicy(req.NoRedirectPolicy())

	// 获取重定向链接
	log.Debugf("发送 GET 请求到 %s", getESurfIpInfoUrl)
	resp, err := client.R().
		Get("http://189.cn/")
	if err != nil {
		log.Error(err)
		return
	}
	location, err := resp.Location()
	if err != nil {
		log.Errorf("获取重定向链接失败：%s", err)
		return
	}
	redirectUrl := location.String()
	log.Debugf("重定向链接：%s", redirectUrl)
	if redirectUrl == getESurfIpInfoUrlRedirect {
		log.Warn("当前已登入")
		err = errors.New("当前已登入")
		return
	}

	// 提取 IP 信息
	match := nasIpRegex.FindStringSubmatch(redirectUrl)
	if len(match) < 2 {
		err = errors.New("获取 nasip 参数失败：未能从重定向链接中提取")
		return
	}
	esurfIpInfo.NasIp = match[1]
	match = clientIpRegex.FindStringSubmatch(redirectUrl)
	if len(match) < 2 {
		err = errors.New("获取 clientip 参数失败：未能从重定向链接中提取")
		return
	}
	esurfIpInfo.ClientIp = match[1]

	return
}
