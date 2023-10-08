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

type getChallengeBody struct {
	NasIP         string `json:"nasip"`
	ClientIP      string `json:"clientip"`
	Mac           string `json:"mac"`
	Username      string `json:"username"`
	Version       string `json:"version"`
	Timestamp     int64  `json:"timestamp"`
	Authenticator string `json:"authenticator"`
}

type GetChallengeResult struct {
	Code      string `json:"rescode"`
	Info      string `json:"resinfo"`
	Challenge string `json:"challenge"`
}

// GetChallenge 获取验证码
func GetChallenge(nasIP, clientIP, mac, username string) (res GetChallengeResult, err error) {

	const getChallengeUrl = urlBase + "/client/vchallenge"

	log.Debugf("获取验证码中...")
	start := time.Now()
	defer func() {
		end := time.Now()
		elapsed := end.Sub(start)
		if err == nil {
			log.Debugf("获取验证码成功，耗时：%s", elapsed)
		} else {
			log.Debugf("获取验证码失败，耗时：%s", elapsed)
		}
	}()

	// 准备请求参数
	var (
		timestamp     = time.Now().Unix()
		authenticator = util.CalMD5Hash(fmt.Sprintf(
			"%s%s%s%s%d%s", version, clientIP, nasIP, mac, timestamp, secret,
		))

		body = getChallengeBody{
			Version:       version,
			Username:      username,
			ClientIP:      clientIP,
			NasIP:         nasIP,
			Mac:           mac,
			Timestamp:     timestamp,
			Authenticator: authenticator,
		}
	)
	log.Debugf("POST body: %+v", body)

	// 发送请求
	log.Debug("发送请求")
	resp, err := req.R().
		SetBody(body).
		SetSuccessResult(&res).
		SetErrorResult(&res).
		Post(getChallengeUrl)

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
		log.Debugf("获取验证码失败（%s）", err)
	}

	log.Debugf("获取验证码成功，验证码为 %s", res.Challenge)
	return
}
