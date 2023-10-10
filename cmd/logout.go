package cmd

import (
	"errors"
	"github.com/Pandaft/go-esurfing/internal/logger"
	"github.com/Pandaft/go-esurfing/internal/util"
	"github.com/Pandaft/go-esurfing/pkg/esurfing"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "登出",
	Long: "登出广东天翼校园网" +
		"\n" +
		"\n必填参数：nasip, clientip",
	Args: func(cmd *cobra.Command, args []string) error {

		log := logger.GetLogger("")

		// 输出参数
		log.Debug("参数：")
		log.Debugf("  - nasip:    %s", nasIP)
		log.Debugf("  - clientip: %s", clientIP)
		log.Debugf("  - mac:      %s", macAddr)
		log.Debugf("  - debug:    %t", &logger.Debug)

		// 验证参数
		if nasIP == "" || clientIP == "" {
			return errors.New("参数 nasip 和 clientip 不能为空")
		}

		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {

		log := logger.GetLogger("")

		// 缺少 mac 参数
		if macAddr == "" {
			log.Warn("缺少 mac 参数，尝试获取中...")
			macAddr = util.GetLocalMACAddr()
			log.Infof("获取成功，mac: %s", macAddr)
		}

		// 登出
		log.Info("请求登出中...")
		res, err := esurfing.Logout(nasIP, clientIP, macAddr)
		if err != nil {
			log.Errorf("登出错误（%s）", err)
			return
		}
		if res.Code != "0" {
			log.Errorf("登出失败（代码：%s  信息：%s）", res.Code, res.Info)
			return
		}
		log.Info("登出成功")

	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	logoutCmd.Flags().SortFlags = false

	logoutCmd.Flags().StringVarP(&nasIP, "nasip", "n", "", "认证服务器 IP")
	logoutCmd.Flags().StringVarP(&clientIP, "clientip", "c", "", "登录设备 IP")
	logoutCmd.Flags().StringVarP(&macAddr, "mac", "m", "", "MAC 地址")
	logoutCmd.Flags().BoolVarP(&logger.Debug, "debug", "d", false, "调试模式")

	_ = logoutCmd.MarkFlagRequired("nasip")
	_ = logoutCmd.MarkFlagRequired("clientip")
}
