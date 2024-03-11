package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sub2sing-box/internal/model"
	. "sub2sing-box/pkg/util"

	"github.com/spf13/cobra"
)

var subscriptions []string
var proxies []string
var template string
var output string
var delete string
var rename map[string]string

func init() {
	convertCmd.Flags().StringSliceVarP(&subscriptions, "subscription", "s", []string{}, "subscription urls")
	convertCmd.Flags().StringSliceVarP(&proxies, "proxy", "p", []string{}, "common proxies")
	convertCmd.Flags().StringVarP(&template, "template", "t", "", "template file path")
	convertCmd.Flags().StringVarP(&output, "output", "o", "", "output file path")
	convertCmd.Flags().StringVarP(&delete, "delete", "d", "", "delete proxy with regex")
	convertCmd.Flags().StringToStringVarP(&rename, "rename", "r", map[string]string{}, "rename proxy with regex")
	RootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Long:  "Convert common proxy to sing-box proxy",
	Short: "Convert common proxy to sing-box proxy",
	Run: func(cmd *cobra.Command, args []string) {
		result := ""
		var err error

		proxyList, err := ConvertSubscriptionsToSProxy(subscriptions)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, proxy := range proxies {
			p, err := ConvertCProxyToSProxy(proxy)
			if err != nil {
				fmt.Println(err)
				return
			}
			proxyList = append(proxyList, p)
		}

		if delete != "" {
			proxyList, err = DeleteProxy(proxyList, delete)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		for k, v := range rename {
			proxyList, err = RenameProxy(proxyList, k, v)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		keep := make(map[int]bool)
		set := make(map[string]struct {
			Proxy model.Proxy
			Count int
		})
		for i, p := range proxyList {
			if _, exists := set[p.Tag]; !exists {
				keep[i] = true
				set[p.Tag] = struct {
					Proxy model.Proxy
					Count int
				}{p, 0}
			} else {
				p1, _ := json.Marshal(p)
				p2, _ := json.Marshal(set[p.Tag])
				if string(p1) != string(p2) {
					set[p.Tag] = struct {
						Proxy model.Proxy
						Count int
					}{p, set[p.Tag].Count + 1}
					keep[i] = true
					proxyList[i].Tag = fmt.Sprintf("%s %d", p.Tag, set[p.Tag].Count)
				} else {
					keep[i] = false
				}
			}
		}
		var newProxyList []model.Proxy
		for i, p := range proxyList {
			if keep[i] {
				newProxyList = append(newProxyList, p)
			}
		}
		proxyList = newProxyList

		if template != "" {
			result, err = MergeTemplate(proxyList, template)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			r, err := json.Marshal(proxyList)
			result = string(r)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		if output != "" {
			err = os.WriteFile(output, []byte(result), 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println(string(result))
		}
	},
}
