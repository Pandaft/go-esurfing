package esurfing

import (
	"errors"
	"fmt"
	"github.com/Pandaft/go-esurfing/internal/util"
	"github.com/charmbracelet/log"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

// loginBody 登入请求体
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

// LoginResult 登入结果
type LoginResult struct {
	Code string `json:"rescode"`
	Info string `json:"resinfo"`
}

// Login 登入
func Login(challenge, nasIP, clientIP, username, password, mac string) (res LoginResult, err error) {

	log.Debugf("登入中...")
	start := time.Now()
	defer func() {
		elapsed := time.Now().Sub(start)
		if err == nil {
			log.Debugf("登入成功，耗时：%s", elapsed)
		} else {
			log.Debugf("登入失败，耗时：%s", elapsed)
		}
	}()

	const loginUrl = urlBase + "/client/login"

	// 准备请求参数
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

	// 检查错误
	if res.Code != "0" {
		err = errors.New(fmt.Sprintf(
			"代码：%s  信息：%s", res.Code, res.Info,
		))
		log.Debugf("登入失败.（%s）", err)
	}

	return
}
