package esurfing

import (
	"fmt"
	"github.com/Pandaft/go-esurfing/internal/util"
	"github.com/charmbracelet/log"
	"github.com/imroc/req/v3"
	"strings"
	"time"
)

// loginBody 登入请求体
type loginBody struct {
	Nasip         string `json:"nasip"`
	Clientip      string `json:"clientip"`
	Mac           string `json:"mac"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Iswifi        string `json:"iswifi"`
	Timestamp     int64  `json:"timestamp"`
	Authenticator string `json:"authenticator"`
}

// LoginResult 登入结果
type LoginResult struct {
	Code string `json:"rescode"`
	Info string `json:"resinfo"`
}

// Login 登入
func Login(nasip, clientip, acc, pwd, mac string) (res LoginResult, err error) {

	const loginUrl = "http://enet.10000.gd.cn:10001/client/login"

	// 获取验证码
	log.Debugf("获取验证码中...")
	challengeRes, err := getChallenge(acc, nasip, clientip)
	if err != nil {
		log.Error(err)
		return
	}
	log.Debugf("获取验证码成功，验证码为 %s", challengeRes.Challenge)

	var (
		timestamp     = time.Now().Unix()
		authenticator = strings.ToUpper(util.CalMD5Hash(fmt.Sprintf(
			"%s%s%s%d%s%s",
			clientip, nasip, mac, timestamp, challengeRes.Challenge, secret,
		)))

		body = loginBody{
			Nasip:         nasip,
			Clientip:      clientip,
			Mac:           mac,
			Username:      acc,
			Password:      pwd,
			Iswifi:        iswifi,
			Timestamp:     timestamp,
			Authenticator: authenticator,
		}
	)

	_, err = req.R().
		SetBody(body).
		SetSuccessResult(&res).
		SetErrorResult(&res).
		Post(loginUrl)
	if err != nil {
		log.Error(err)
	}
	return
}
