package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sub2sing-box/model"
	"sub2sing-box/parser"
	. "sub2sing-box/util"

	"github.com/spf13/cobra"
)

func Url(cmd *cobra.Command, args []string) {
	proxyList := make([]model.Proxy, 0)
	if cmd.Flag("url").Changed {
		urls, _ := cmd.Flags().GetStringSlice("url")
		for _, url := range urls {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			proxy, err := DecodeBase64(string(data))
			if err != nil {
				fmt.Println(err)
				return
			}
			proxies := strings.Split(proxy, "\n")
			for _, p := range proxies {
				if strings.HasPrefix(p, "ss://") {
					proxy, err := parser.ParseShadowsocks(p)
					if err != nil {
						fmt.Println(proxy)
					}
					proxyList = append(proxyList, proxy)
				} else if strings.HasPrefix(p, "vmess://") {
					proxy, err := parser.ParseVmess(p)
					if err != nil {
						fmt.Println(proxy)
					}
					proxyList = append(proxyList, proxy)
				} else if strings.HasPrefix(p, "trojan://") {
					proxy, err := parser.ParseTrojan(p)
					if err != nil {
						fmt.Println(proxy)
					}
					proxyList = append(proxyList, proxy)
				} else if strings.HasPrefix(p, "vless://") {
					proxy, err := parser.ParseVless(p)
					if err != nil {
						fmt.Println(proxy)
					}
					proxyList = append(proxyList, proxy)
				} else if strings.HasPrefix(p, "hysteria://") {
					proxy, err := parser.ParseHysteria(p)
					if err != nil {
						fmt.Println(proxy)
					}
					proxyList = append(proxyList, proxy)
				} else if strings.HasPrefix(p, "hy2://") || strings.HasPrefix(p, "hysteria2://") {
					proxy, err := parser.ParseHysteria2(p)
					if err != nil {
						fmt.Println(proxy)
					}
					proxyList = append(proxyList, proxy)
				}
			}
		}
		result, err := json.Marshal(proxyList)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(string(result))
		}
	} else {
		fmt.Println("No URLs provided")
	}
}
