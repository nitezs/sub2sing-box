package cmd

import (
	"log/slog"
	"sub2sing-box/common"
	"sub2sing-box/model"

	"github.com/spf13/cobra"
)

var ParseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parse other kernel config to sing-box config",
	Long:  "parse other kernel config to sing-box config",
	Run:   parseFn,
}

func init() {
	ParseCmd.Flags().StringSliceP("general-config", "g", nil, "The path or URL of the general configuration to be converted")
	ParseCmd.Flags().StringSliceP("clash-config", "c", nil, "The path or URL of the clash configuration to be converted")
}

func parseFn(cmd *cobra.Command, args []string) {
	// generalConfigs, err := cmd.Flags().GetStringSlice("general-config")
	clashConfig, err := cmd.Flags().GetStringSlice("clash-config")
	var singBox *model.SingBox = &model.SingBox{}

	if err != nil {
		slog.Error("get flags failed", err)
	}

	for index, config := range clashConfig {
		configData, err := common.GetConfig(config)
		if err != nil {
			slog.Error("get config failed", err)
		}
		// parse proxies
		common.ParseClashProxies(singBox, configData)

		// parse rules
		if index == 0 {
			common.ParseClashRules(singBox, configData)
		}
	}
	// add all proxies selector
	common.AddAllProxiesSelector(singBox)
	// add country selectors

}
