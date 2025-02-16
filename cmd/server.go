package cmd

import (
	"github.com/nitezs/sub2sing-box/api"

	"github.com/spf13/cobra"
)

var (
	port uint16
	bind string
)

func init() {
	runCmd.Flags().Uint16VarP(&port, "port", "p", 8080, "server port")
	runCmd.Flags().StringVarP(&bind, "bind", "b", "0.0.0.0", "bind address (e.g., 0.0.0.0, 127.0.0.1, localhost)")
	RootCmd.AddCommand(runCmd)
}

+var runCmd = &cobra.Command{
	Use:   "server",
	Long:  "Run the server",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		api.RunServer(bind, port)
	},
}