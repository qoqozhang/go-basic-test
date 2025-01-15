package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "server-cli",
	Short:   "server-cli",
	Long:    "A server-cli tools",
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	TraverseChildren: true,
}

func tip() {
	fmt.Printf("%s\n", "欢迎使用 server-cli 工具，可以使用 -h 查看命令")
	fmt.Printf("%s\n", "也可以参考https://xxx 的相关内容")
}

func Execute() {
	rootCmd.AddCommand(addCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
