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

const getChallengeUrl = urlBase + "/client/vchallenge"

// GetChallenge 获取验证码
func GetChallenge(nasIP, clientIP, mac, username string) (res GetChallengeResult, err error) {

	log := logger.GetLogger("获取验证码")

	// 输出耗时
	defer util.MeasureExecTime(time.Now(), log)

	// 准备参数
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

	// 调试输出
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

	// 判断结果
	if res.Code != "0" {
		err = errors.New(fmt.Sprintf(
			"代码：%s  信息：%s", res.Code, res.Info,
		))
		log.Debugf("获取失败：%s", err)
	} else {
		log.Debugf("获取成功，验证码为 %s", res.Challenge)
	}

	return
}
