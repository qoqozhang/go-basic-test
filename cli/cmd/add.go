package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	addUser string
	addCmd  = &cobra.Command{
		Use:   "add",
		Short: "add command",
		Long:  "add command long comment",
		Run: func(cmd *cobra.Command, args []string) {
			if addUser == "" {
				cmd.Usage()
			} else {
				fmt.Printf("add user : %s\n", addUser)
			}
		},
	}
	addPassword = &cobra.Command{
		Use:   "password",
		Short: "add user's password",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				cmd.Usage()
			} else {
				fmt.Printf("password: %s\n", args[0])
			}
		},
	}
)

func init() {
	addCmd.Flags().StringVarP(&addUser, "user", "u", "", "add user name parameter")
	addCmd.AddCommand(addPassword)
}
