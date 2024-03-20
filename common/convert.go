package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	model2 "sub2sing-box/model"
	"sub2sing-box/parser"
	"sub2sing-box/util"
)

func Convert(
	subscriptions []string,
	proxies []string,
	template string,
	delete string,
	rename map[string]string,
	group bool,
	groupType string,
	sortKey string,
	sortType string,
) (string, error) {
	result := ""
	var err error

	proxyList, err := ConvertSubscriptionsToSProxy(subscriptions)
	if err != nil {
		return "", err
	}
	for _, proxy := range proxies {
		p, err := ConvertCProxyToSProxy(proxy)
		if err != nil {
			return "", err
		}
		proxyList = append(proxyList, p)
	}

	if delete != "" {
		proxyList, err = DeleteProxy(proxyList, delete)
		if err != nil {
			return "", err
		}
	}

	for k, v := range rename {
		proxyList, err = RenameProxy(proxyList, k, v)
		if err != nil {
			return "", err
		}
	}

	keep := make(map[int]bool)
	set := make(map[string]struct {
		Proxy model2.Proxy
		Count int
	})
	for i, p := range proxyList {
		if _, exists := set[p.Tag]; !exists {
			keep[i] = true
			set[p.Tag] = struct {
				Proxy model2.Proxy
				Count int
			}{p, 0}
		} else {
			p1, _ := json.Marshal(p)
			p2, _ := json.Marshal(set[p.Tag])
			if string(p1) != string(p2) {
				set[p.Tag] = struct {
					Proxy model2.Proxy
					Count int
				}{p, set[p.Tag].Count + 1}
				keep[i] = true
				proxyList[i].Tag = fmt.Sprintf("%s %d", p.Tag, set[p.Tag].Count)
			} else {
				keep[i] = false
			}
		}
	}
	var newProxyList []model2.Proxy
	for i, p := range proxyList {
		if keep[i] {
			newProxyList = append(newProxyList, p)
		}
	}
	proxyList = newProxyList
	var outbounds []model2.Outbound
	ps, err := json.Marshal(&proxyList)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(ps, &outbounds)
	if err != nil {
		return "", err
	}
	if group {
		outbounds = AddCountryGroup(outbounds, groupType, sortKey, sortType)
	}
	if template != "" {
		result, err = MergeTemplate(outbounds, template)
		if err != nil {
			return "", err
		}
	} else {
		r, err := json.Marshal(outbounds)
		result = string(r)
		if err != nil {
			return "", err
		}
	}

	return string(result), nil
}

func AddCountryGroup(proxies []model2.Outbound, groupType string, sortKey string, sortType string) []model2.Outbound {
	newGroup := make(map[string]model2.Outbound)
	for _, p := range proxies {
		if p.Type != "selector" && p.Type != "urltest" {
			country := model2.GetContryName(p.Tag)
			if group, ok := newGroup[country]; ok {
				group.Outbounds = append(group.Outbounds, p.Tag)
				newGroup[country] = group
			} else {
				newGroup[country] = model2.Outbound{
					Tag:                       country,
					Type:                      groupType,
					Outbounds:                 []string{p.Tag},
					InterruptExistConnections: true,
				}
			}
		}
	}
	var groups []model2.Outbound
	for _, p := range newGroup {
		groups = append(groups, p)
	}
	if sortType == "asc" {
		switch sortKey {
		case "tag":
			sort.Sort(model2.SortByTag(groups))
		case "num":
			sort.Sort(model2.SortByNumber(groups))
		default:
			sort.Sort(model2.SortByTag(groups))
		}
	} else {
		switch sortKey {
		case "tag":
			sort.Sort(sort.Reverse(model2.SortByTag(groups)))
		case "num":
			sort.Sort(sort.Reverse(model2.SortByNumber(groups)))
		default:
			sort.Sort(sort.Reverse(model2.SortByTag(groups)))
		}
	}
	return append(proxies, groups...)
}

func MergeTemplate(outbounds []model2.Outbound, template string) (string, error) {
	var config model2.Config
	var err error
	if strings.HasPrefix(template, "http") {
		data, err := util.Fetch(template, 3)
		if err != nil {
			return "", err
		}
		err = json.Unmarshal([]byte(data), &config)
		if err != nil {
			return "", err
		}
	} else {
		if !strings.Contains(template, string(filepath.Separator)) {
			if _, err := os.Stat(template); os.IsNotExist(err) {
				template = filepath.Join("templates", template)
			}
		}
		config, err = ReadTemplate(template)
	}
	proxyTags := make([]string, 0)
	groupTags := make([]string, 0)
	groups := make(map[string]model2.Outbound)
	if err != nil {
		return "", err
	}
	for _, p := range outbounds {
		if model2.IsCountryGroup(p.Tag) {
			groupTags = append(groupTags, p.Tag)
			reg := regexp.MustCompile("[A-Za-z]{2}")
			country := reg.FindString(p.Tag)
			groups[country] = p
		} else {
			proxyTags = append(proxyTags, p.Tag)
		}
	}
	reg := regexp.MustCompile("<[A-Za-z]{2}>")
	for i, outbound := range config.Outbounds {
		var parsedOutbound []string = make([]string, 0)
		for _, o := range outbound.Outbounds {
			if o == "<all-proxy-tags>" {
				parsedOutbound = append(parsedOutbound, proxyTags...)
			} else if o == "<all-country-tags>" {
				parsedOutbound = append(parsedOutbound, groupTags...)
			} else if reg.MatchString(o) {
				country := strings.ToUpper(strings.Trim(reg.FindString(o), "<>"))
				if group, ok := groups[country]; ok {
					parsedOutbound = append(parsedOutbound, group.Outbounds...)
				}
			} else {
				parsedOutbound = append(parsedOutbound, o)
			}
		}
		config.Outbounds[i].Outbounds = parsedOutbound
	}
	config.Outbounds = append(config.Outbounds, outbounds...)
	data, err := json.Marshal(config)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ConvertCProxyToSProxy(proxy string) (model2.Proxy, error) {
	for prefix, parseFunc := range parser.ParserMap {
		if strings.HasPrefix(proxy, prefix) {
			proxy, err := parseFunc(proxy)
			if err != nil {
				return model2.Proxy{}, err
			}
			return proxy, nil
		}
	}
	return model2.Proxy{}, errors.New("Unknown proxy format")
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

func ConvertSubscriptionsToSProxy(urls []string) ([]model2.Proxy, error) {
	proxyList := make([]model2.Proxy, 0)
	for _, url := range urls {
		data, err := util.Fetch(url, 3)
		if err != nil {
			return nil, err
		}
		proxy, err := util.DecodeBase64(data)
		if err != nil {
			return nil, err
		}
		proxies := strings.Split(proxy, "\n")
		for _, p := range proxies {
			for prefix, parseFunc := range parser.ParserMap {
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

func ReadTemplate(path string) (model2.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return model2.Config{}, err
	}
	var res model2.Config
	err = json.Unmarshal(data, &res)
	if err != nil {
		return model2.Config{}, err
	}
	return res, nil
}

func DeleteProxy(proxies []model2.Proxy, regex string) ([]model2.Proxy, error) {
	reg, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}
	var newProxies []model2.Proxy
	for _, p := range proxies {
		if !reg.MatchString(p.Tag) {
			newProxies = append(newProxies, p)
		}
	}
	return newProxies, nil
}

func RenameProxy(proxies []model2.Proxy, regex string, replaceText string) ([]model2.Proxy, error) {
	reg, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}
	for i, p := range proxies {
		if reg.MatchString(p.Tag) {
			proxies[i].Tag = reg.ReplaceAllString(p.Tag, replaceText)
		}
	}
	return proxies, nil
}
