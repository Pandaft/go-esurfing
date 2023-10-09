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

type loginBody struct {
	NasIP         string `json:"nasip"`
	ClientIP      string `json:"clientip"`
	Mac           string `json:"mac"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	IsWiFi        string `json:"iswifi"`
	Timestamp     int64  `json:"timestamp"`
	Authenticator string `json:"authenticator"`
}

type LoginResult struct {
	Code string `json:"rescode"`
	Info string `json:"resinfo"`
}

const loginUrl = urlBase + "/client/login"

// Login 登入
func Login(challenge, nasIP, clientIP, username, password, mac string) (res LoginResult, err error) {

	log := logger.GetLogger("登入校园网")

	// 执行耗时
	start := time.Now()
	log.Debug("开始")
	defer func() {
		elapsed := time.Now().Sub(start)
		log.Debugf("结束，耗时：%s", elapsed)
	}()

	// 准备参数
	var (
		timestamp     = time.Now().Unix()
		authenticator = util.CalMD5Hash(fmt.Sprintf(
			"%s%s%s%d%s%s", clientIP, nasIP, mac, timestamp, challenge, secret,
		))

		body = loginBody{
			NasIP:         nasIP,
			ClientIP:      clientIP,
			Mac:           mac,
			Username:      username,
			Password:      password,
			IsWiFi:        iswifi,
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
		Post(loginUrl)

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
		log.Debugf("登入失败：%s", err)
	} else {
		log.Debugf("登入成功")
	}

	return
}
