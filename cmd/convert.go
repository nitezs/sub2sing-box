package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sub2sing-box/constant"
	"sub2sing-box/model"
	. "sub2sing-box/util"

	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Long:  "Convert common proxy to sing-box proxy",
	Short: "Convert common proxy to sing-box proxy",
	Run: func(cmd *cobra.Command, args []string) {
		subscriptions, _ := cmd.Flags().GetStringSlice("subscription")
		proxies, _ := cmd.Flags().GetStringSlice("proxy")
		template, _ := cmd.Flags().GetString("template")
		output, _ := cmd.Flags().GetString("output")
		if template == "" {
			proxyList, err := ConvertSubscriptionsToSProxy(subscriptions)
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, p := range proxies {
				result, err := ConvertCProxyToSProxy(p)
				if err != nil {
					fmt.Println(err)
					return
				}
				proxyList = append(proxyList, result)
			}
			result, err := json.Marshal(proxyList)
			if err != nil {
				fmt.Println(err)
				return
			}
			if output != "" {
				err = os.WriteFile(output, result, 0666)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				fmt.Println(string(result))
			}
		} else {
			config, err := ConvertWithTemplate(subscriptions, proxies, template)
			if err != nil {
				fmt.Println(err)
				return
			}
			data, err := json.Marshal(config)
			if err != nil {
				fmt.Println(err)
				return
			}
			if output != "" {
				err = os.WriteFile(output, data, 0666)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				fmt.Println(string(data))
			}
		}
	},
}

func init() {
	convertCmd.Flags().StringSliceP("subscription", "s", []string{}, "subscription urls")
	convertCmd.Flags().StringSliceP("proxy", "p", []string{}, "common proxies")
	convertCmd.Flags().StringP("template", "t", "", "path of template file")
	convertCmd.Flags().StringP("output", "o", "", "output file path")
	RootCmd.AddCommand(convertCmd)
}

func Convert(urls []string, proxies []string) ([]model.Proxy, error) {
	proxyList := make([]model.Proxy, 0)
	newProxies, err := ConvertSubscriptionsToSProxy(urls)
	if err != nil {
		return nil, err
	}
	proxyList = append(proxyList, newProxies...)
	for _, p := range proxies {
		proxy, err := ConvertCProxyToSProxy(p)
		if err != nil {
			return nil, err
		}
		proxyList = append(proxyList, proxy)
	}
	return proxyList, nil
}

func ConvertWithTemplate(urls []string, proxies []string, template string) (model.Config, error) {
	proxyList := make([]model.Proxy, 0)
	newProxies, err := ConvertSubscriptionsToSProxy(urls)
	newOutboundTagList := make([]string, 0)
	if err != nil {
		return model.Config{}, err
	}
	proxyList = append(proxyList, newProxies...)
	for _, p := range proxies {
		proxy, err := ConvertCProxyToSProxy(p)
		if err != nil {
			return model.Config{}, err
		}
		proxyList = append(proxyList, proxy)
	}
	config, err := ReadTemplate(template)
	if err != nil {
		return model.Config{}, err
	}
	ps, err := json.Marshal(proxyList)
	if err != nil {
		return model.Config{}, err
	}
	var newOutbounds []model.Outbound
	err = json.Unmarshal(ps, &newOutbounds)
	if err != nil {
		return model.Config{}, err
	}
	for _, outbound := range newOutbounds {
		newOutboundTagList = append(newOutboundTagList, outbound.Tag)
	}
	config.Outbounds = append(config.Outbounds, newOutbounds...)
	for i, outbound := range config.Outbounds {
		if outbound.Type == "urltest" || outbound.Type == "selector" {
			var parsedOutbound []string = make([]string, 0)
			for _, o := range outbound.Outbounds {
				if o == "<all-proxy-tags>" {
					parsedOutbound = append(parsedOutbound, newOutboundTagList...)
				} else {
					parsedOutbound = append(parsedOutbound, o)
				}
			}
			config.Outbounds[i].Outbounds = parsedOutbound
		}
	}
	return config, nil
}

func ConvertCProxyToSProxy(proxy string) (model.Proxy, error) {
	for prefix, parseFunc := range constant.ParserMap {
		if strings.HasPrefix(proxy, prefix) {
			proxy, err := parseFunc(proxy)
			if err != nil {
				return model.Proxy{}, err
			}
			return proxy, nil
		}
	}
	return model.Proxy{}, errors.New("Unknown proxy format")
}

func ConvertCProxyToJson(proxy string) (string, error) {
	sProxy, err := ConvertCProxyToSProxy(proxy)
	if err != nil {
		return "", err
	}
	data, err := json.Marshal(&sProxy)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func FetchSubscription(url string, maxRetryTime int) (string, error) {
	retryTime := 0
	var err error
	for retryTime < maxRetryTime {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			retryTime++
			continue
		}
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			retryTime++
			continue
		}
		return string(data), err
	}
	return "", err
}

func ConvertSubscriptionsToSProxy(urls []string) ([]model.Proxy, error) {
	proxyList := make([]model.Proxy, 0)
	for _, url := range urls {
		data, err := FetchSubscription(url, 3)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		proxy, err := DecodeBase64(data)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		proxies := strings.Split(proxy, "\n")
		for _, p := range proxies {
			for prefix, parseFunc := range constant.ParserMap {
				if strings.HasPrefix(p, prefix) {
					proxy, err := parseFunc(p)
					if err != nil {
						return nil, err
					}
					proxyList = append(proxyList, proxy)
				}
			}
		}
	}
	return proxyList, nil
}

func ConvertSubscriptionsToJson(urls []string) (string, error) {
	proxyList, err := ConvertSubscriptionsToSProxy(urls)
	if err != nil {
		return "", err
	}
	result, err := json.Marshal(proxyList)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func ReadTemplate(path string) (model.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return model.Config{}, err
	}
	var res model.Config
	err = json.Unmarshal(data, &res)
	if err != nil {
		return model.Config{}, err
	}
	return res, nil
}
