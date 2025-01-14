package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	version    string = "v0.0.1"
	versionCmd        = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: "server-cli version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}
)
