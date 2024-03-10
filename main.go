package main

import (
	"fmt"
	"sub2sing-box/cmd"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "process",
	Run: cmd.Url,
}

func init() {
	rootCmd.Flags().StringSliceP("url", "u", []string{}, "URLs to process")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
