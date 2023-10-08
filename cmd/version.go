package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const version = "v0.1.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "输出版本",
	Long:  "输出当前 go-esurfing 具体版本",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
