package main

import (
	"fmt"
	"sub2sing-box/cmd"
)

var Version string

func init() {
	Version = "dev"
	cmd.SetVersion(Version)
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
