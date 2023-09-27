package esurfing

import (
	"errors"
	"fmt"
	"github.com/Pandaft/go-esurfing/internal/util"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

type getChallengeBody struct {
	Nasip         string `json:"nasip"`
	Clientip      string `json:"clientip"`
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

// getChallenge 获取验证码
func getChallenge(username, nasIp, clientIp string) (res GetChallengeResult, err error) {

	const getChallengeUrl = "http://enet.10000.gd.cn:10001/client/vchallenge"

	var (
		mac           = util.GetLocalMACAddr()
		timestamp     = time.Now().Unix()
		authenticator = strings.ToUpper(util.CalMD5Hash(fmt.Sprintf(
			"%s%s%s%s%d%s", version, clientIp, nasIp, mac, timestamp, secret,
		)))

		body = getChallengeBody{
			Version:  version,
			Username: username,
			Clientip: clientIp,
			Nasip:    nasIp,

			Mac:           mac,
			Timestamp:     timestamp,
			Authenticator: authenticator,
		}
	)

	_, err = req.R().
		SetBody(body).
		SetSuccessResult(&res).
		SetErrorResult(&res).
		Post(getChallengeUrl)

	if err != nil {
		return
	}

	if res.Code != "0" {
		err = errors.New(fmt.Sprintf(
			"错误代码：%s  错误信息：%s", res.Code, res.Info,
		))
	}

	return
}
