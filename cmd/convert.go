package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nitezs/sub2sing-box/common"
	"github.com/nitezs/sub2sing-box/model"

	"github.com/spf13/cobra"
)

var (
	subscriptions []string
	proxies       []string
	template      string
	output        string
	delete        string
	rename        map[string]string
	group         bool
	groupType     string
	sortKey       string
	sortType      string
	config        string
)

func init() {
	convertCmd.Flags().StringSliceVarP(&subscriptions, "subscription", "s", make([]string, 0), "subscription URLs")
	convertCmd.Flags().StringSliceVarP(&proxies, "proxy", "p", make([]string, 0), "proxies share links")
	convertCmd.Flags().StringVarP(&template, "template", "t", "", "template file path or URL")
	convertCmd.Flags().StringVarP(&output, "output", "o", "", "output file path")
	convertCmd.Flags().StringVarP(&delete, "delete", "d", "", "delete proxies with regex")
	convertCmd.Flags().StringToStringVarP(&rename, "rename", "r", make(map[string]string), "rename proxies with regex")
	convertCmd.Flags().BoolVarP(&group, "group", "g", false, "grouping proxies by country")
	convertCmd.Flags().StringVarP(&groupType, "group-type", "G", "", "group type, selector or urltest")
	convertCmd.Flags().StringVarP(&sortKey, "sort", "S", "", "sort key, tag or num")
	convertCmd.Flags().StringVarP(&sortType, "sort-type", "T", "", "sort type, asc or desc")
	convertCmd.Flags().StringVarP(&config, "config", "c", "", "configuration file path")
	RootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Long:  "Convert common proxy to sing-box proxy",
	Short: "Convert common proxy to sing-box proxy",
	Run:   convertRun,
}

func convertRun(cmd *cobra.Command, args []string) {
	loadConfig()
	result, err := common.Convert(
		subscriptions,
		proxies,
		template,
		delete,
		rename,
		group,
		groupType,
		sortKey,
		sortType,
	)
	if err != nil {
		fmt.Println("Conversion error:", err)
		return
	}
	if output != "" {
		err = os.WriteFile(output, []byte(result), 0666)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Println("Config has been saved in:", output)
	} else {
		fmt.Println(result)
	}
}

func loadConfig() {
	if config == "" {
		if wd, err := os.Getwd(); err == nil {
			config = filepath.Join(wd, "sub2sing-box.json")
			if _, err := os.Stat(config); os.IsNotExist(err) {
				return
			}
		} else {
			fmt.Println("Error getting working directory:", err)
			return
		}
	}

	bytes, err := os.ReadFile(config)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	var cfg model.ConvertRequest
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		fmt.Println("Error parsing config JSON:", err)
		return
	}

	mergeConfig(cfg)
}

func mergeConfig(cfg model.ConvertRequest) {
	subscriptions = append(subscriptions, cfg.Subscriptions...)
	proxies = append(proxies, cfg.Proxies...)
	if template == "" {
		template = cfg.Template
	}
	if delete == "" {
		delete = cfg.Delete
	}
	for k, v := range cfg.Rename {
		rename[k] = v
	}
	if groupType == "" {
		if cfg.GroupType != "" {
			groupType = cfg.GroupType
		} else {
			groupType = "selector"
		}
	}
	if sortKey == "" {
		if cfg.SortKey != "" {
			sortKey = cfg.SortKey
		} else {
			sortKey = "tag"
		}
	}
	if sortType == "" {
		if cfg.SortType != "" {
			sortType = cfg.SortType
		} else {
			sortType = "asc"
		}
	}
	if output == "" {
		output = cfg.Output
	}
}
