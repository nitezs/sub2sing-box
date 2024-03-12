package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{}

func SetVersion(version string) {
	RootCmd.Version = version
}
