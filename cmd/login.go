package cmd

import (
	"errors"
	"github.com/Pandaft/go-esurfing/internal/logger"
	"github.com/Pandaft/go-esurfing/internal/util"
	"github.com/Pandaft/go-esurfing/pkg/esurfing"
	"github.com/spf13/cobra"
	"time"
)

var (
	nasIP    string
	clientIP string
	macAddr  string
	username string
	password string
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "登入",
	Long: "登入广东天翼校园网" +
		"\n" +
		"\n必填参数：username, password" +
		"\n" +
		"\n对于 nasip 和 clientip 参数：" +
		"\n  - 本机未登入，且在本机登入时，可不填写" +
		"\n  - 本机已登入，或在远程登入时，必须填写" +
		"\n" +
		"\n对于 mac 参数：" +
		"\n  - 暂未发现对登入功能有实际影响" +
		"\n  - 不填写默认为 00-00-00-00-00-00",
	Args: func(cmd *cobra.Command, args []string) error {

		log := logger.GetLogger("")

		// 输出参数
		log.Debug("参数：")
		log.Debugf("  - nasip:    %s", nasIP)
		log.Debugf("  - clientip: %s", clientIP)
		log.Debugf("  - mac:      %s", macAddr)
		log.Debugf("  - username: %s", username)
		log.Debugf("  - password: %s", password)
		log.Debugf("  - debug:    %t", logger.Debug)

		// 验证参数
		if username == "" || password == "" {
			return errors.New("账号或密码不能为空")
		}

		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {

		log := logger.GetLogger("")

		// 输出版本
		log.Infof("版本：%s", version)

		// 输出耗时
		defer util.MeasureExecTime(time.Now(), log)

		// 缺少 nasip 或 clientip 参数
		if nasIP == "" || clientIP == "" {
			log.Warn("缺少 nasip 或 clientip 参数，尝试获取中...")
			esurfIpInfo, err := esurfing.GetESurfIpInfo()
			if err != nil {
				log.Errorf("获取失败：%s", err)
				return
			}
			nasIP, clientIP = esurfIpInfo.NasIp, esurfIpInfo.ClientIp
			log.Infof("获取成功（nasip: %s  clientip: %s）", nasIP, clientIP)
		}

		// 缺少 mac 参数
		if macAddr == "" {
			log.Warn("缺少 mac 参数，尝试获取中...")
			macAddr = util.GetLocalMACAddr()
			log.Infof("获取成功，mac: %s", macAddr)
		}

		// 获取验证码
		log.Info("获取验证码中...")
		challengeRes, err := esurfing.GetChallenge(nasIP, clientIP, macAddr, username)
		if err != nil {
			log.Errorf("获取验证码失败（%s）", err)
			return
		}
		log.Infof("获取验证码成功（验证码：%s）", challengeRes.Challenge)

		// 登入
		log.Info("请求登入中...")
		res, err := esurfing.Login(challengeRes.Challenge, nasIP, clientIP, username, password, macAddr)
		if err != nil {
			log.Errorf("登入错误（%s）", err)
			return
		}
		if res.Code != "0" {
			log.Errorf("登入失败（代码：%s  信息：%s）", res.Code, res.Info)
			return
		}
		log.Info("登入成功")

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().SortFlags = false

	loginCmd.Flags().StringVarP(&nasIP, "nasip", "n", "", "认证服务器 IP")
	loginCmd.Flags().StringVarP(&clientIP, "clientip", "c", "", "登录设备 IP")
	loginCmd.Flags().StringVarP(&macAddr, "mac", "m", "", "MAC 地址")
	loginCmd.Flags().StringVarP(&username, "username", "u", "", "账号")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "密码")
	loginCmd.Flags().BoolVarP(&logger.Debug, "debug", "d", false, "调试模式")

	_ = loginCmd.MarkFlagRequired("username")
	_ = loginCmd.MarkFlagRequired("password")
}
