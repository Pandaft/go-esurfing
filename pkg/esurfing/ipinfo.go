package esurfing

import (
	"errors"
	"github.com/charmbracelet/log"
	"github.com/imroc/req/v3"
	"regexp"
	"time"
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

	log.Debugf("获取 IP 信息中...")
	start := time.Now()
	defer func() {
		elapsed := time.Now().Sub(start)
		if err == nil {
			log.Debugf("获取 IP 信息成功，耗时：%s", elapsed)
		} else {
			log.Debugf("获取 IP 信息失败，耗时：%s", elapsed)
		}
	}()

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
		log.Debugf("获取重定向链接失败：%s", err)
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
