package main

import (
	"fmt"

	"github.com/nitezs/sub2sing-box/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
