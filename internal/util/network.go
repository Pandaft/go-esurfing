package util

import (
	"github.com/charmbracelet/log"
)

// GetLocalMACAddr 获取本机 MAC 地址
func GetLocalMACAddr() string {
	const macAddr = "00-00-00-00-00-00"
	log.Warnf("暂未发现 MAC 地址对登录有实际影响，因此 MAC 参数将被设定为 %s", macAddr)
	log.Warnf("如果有实际影响，需手动设定 MAC 参数，必要时请到 GitHub 提 issue 反馈")
	return macAddr
}
