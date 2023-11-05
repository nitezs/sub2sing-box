package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sub2sing-box",
	Short: "a tool to generate sing-box config",
	Long:  "a tool to generate sing-box config",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(ParseCmd)
}

func MyCommand() *cobra.Command {
	return rootCmd
}
