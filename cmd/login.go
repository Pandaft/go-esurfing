package cmd

import (
	"errors"
	"github.com/Pandaft/go-esurfing/internal/util"
	"github.com/Pandaft/go-esurfing/pkg/esurfing"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var (
	nasIP    string
	clientIp string
	macAddr  string
	pwd      string
	acc      string
	debug    bool
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "登入",
	Long: "登入广东天翼校园网" +
		"\n" +
		"\n对于 nasip 和 clientip 参数：" +
		"\n  - 本机未登入，且在本机登入时，可不填写" +
		"\n  - 本机已登入，或在远程登入时，必须填写",
	Args: func(cmd *cobra.Command, args []string) error {
		// 调试模式
		if debug {
			log.SetLevel(log.DebugLevel)
		}

		// 输出参数
		log.Debug("参数：")
		log.Debugf("  - 账号：%s", acc)
		log.Debugf("  - 密码：%s", pwd)
		log.Debugf("  - nasIP：%s", nasIP)
		log.Debugf("  - clientIP：%s", clientIp)
		log.Debugf("  - 调试：%t", debug)

		// 验证参数
		if acc == "" || pwd == "" {
			return errors.New("账号或密码不能为空")
		}

		// 缺少 nasip 或 clientip 参数
		if nasIP == "" || clientIp == "" {
			log.Warn("缺少 nasip 或 clientip 参数，尝试获取中...")
			esurfIpInfo, err := esurfing.GetESurfIpInfo()
			if err != nil {
				log.Errorf("获取失败：%s", err)
				return err
			}
			nasIP, clientIp = esurfIpInfo.NasIp, esurfIpInfo.ClientIp
			log.Infof("获取成功：nasip=%s  clientip=%s", nasIP, clientIp)
		}

		// 缺省 MAC 地址
		if macAddr == "" {
			macAddr = util.GetLocalMACAddr()
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		res, err := esurfing.Login(nasIP, clientIp, acc, pwd, macAddr)
		if err != nil {
			log.Error(err)
			return
		}
		if res.Code == "0" {
			log.Info("登入成功")
		} else {
			log.Errorf("登入失败，代码：%s  信息：%s", res.Code, res.Info)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&acc, "acc", "a", "", "账号")
	loginCmd.Flags().StringVarP(&pwd, "pwd", "p", "", "密码")
	loginCmd.Flags().StringVarP(&nasIP, "nasip", "n", "", "nasIP")
	loginCmd.Flags().StringVarP(&clientIp, "clientip", "c", "", "clientIP")
	loginCmd.Flags().StringVarP(&macAddr, "mac", "m", "", "MAC 地址")
	loginCmd.Flags().BoolVarP(&debug, "debug", "d", true, "调试模式")

	_ = loginCmd.MarkFlagRequired("acc")
	_ = loginCmd.MarkFlagRequired("pwd")
}
