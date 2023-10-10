package esurfing

import (
	"errors"
	"fmt"
	"github.com/Pandaft/go-esurfing/internal/logger"
	"github.com/Pandaft/go-esurfing/internal/util"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

type logoutBody struct {
	NasIP         string `json:"nasip"`
	ClientIP      string `json:"clientip"`
	Mac           string `json:"mac"`
	Secret        string `json:"secret"`
	Timestamp     int64  `json:"timestamp"`
	Authenticator string `json:"authenticator"`
}

type LogoutResult struct {
	Code string `json:"rescode"`
	Info string `json:"resinfo"`
}

const logoutUrl = urlBase + "/client/logout"

// Logout 登出
func Logout(nasIP, clientIP, mac string) (res LogoutResult, err error) {

	log := logger.GetLogger("登出校园网")

	// 执行耗时
	defer util.MeasureExecTime(time.Now(), log)

	// 准备参数
	var (
		timestamp     = time.Now().Unix()
		authenticator = util.CalMD5Hash(fmt.Sprintf(
			"%s%s%s%d%s", clientIP, nasIP, mac, timestamp, secret,
		))

		body = logoutBody{
			NasIP:         nasIP,
			ClientIP:      clientIP,
			Mac:           mac,
			Secret:        secret,
			Timestamp:     timestamp,
			Authenticator: authenticator,
		}
	)

	// 调试输出
	log.Debugf("POST body: %+v", body)

	// 发送请求
	log.Debug("发送请求")
	resp, err := req.R().
		SetBody(body).
		SetSuccessResult(&res).
		SetErrorResult(&res).
		Post(logoutUrl)

	// 发生错误
	if err != nil {
		log.Debugf("发生错误：%s", err)
		return
	}

	// 调试输出
	log.Debugf("resp body: %s", strings.Trim(resp.String(), "\n"))
	log.Debugf("res: %+v", res)

	// 判断结果
	if res.Code != "0" {
		err = errors.New(fmt.Sprintf(
			"代码：%s  信息：%s", res.Code, res.Info,
		))
		log.Debugf("登出失败：%s", err)
	} else {
		log.Debugf("登出成功")
	}

	return
}
