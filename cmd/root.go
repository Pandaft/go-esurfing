package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"os"
	"time"
)

const (
	shortDesc = "基于 Go 实现登入和登出广东天翼校园网的命令行工具"
	githubUrl = "https://github.com/Pandaft/go-esurfing"
)

var rootCmd = &cobra.Command{
	Use:   "go-esurfing",
	Short: shortDesc,
	Long:  fmt.Sprintf("%s (%s)\nGitHub: %s", shortDesc, version, githubUrl),
	Run: func(cmd *cobra.Command, args []string) {
		log.Warn("此程序为命令行工具，请带命令和参数运行，程序将在 3 秒后自动关闭...")
		time.Sleep(time.Second * 3)
		os.Exit(1)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
