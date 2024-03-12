package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: " + RootCmd.Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
