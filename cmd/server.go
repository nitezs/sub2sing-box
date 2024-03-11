package cmd

import (
	"sub2sing-box/api"

	"github.com/spf13/cobra"
)

var port uint16

func init() {
	runCmd.Flags().Uint16VarP(&port, "port", "p", 8080, "server port")
	RootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "server",
	Long:  "Run the server",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		api.RunServer(port)
	},
}
