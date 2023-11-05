package main

import (
	"sub2sing-box/cmd"
	"sub2sing-box/log"

	"github.com/spf13/cobra"
)

var Cmd *cobra.Command

func init() {
	Cmd = cmd.MyCommand()
	log.SetLogger("debug")
}

func main() {
	Cmd.Execute()
}
