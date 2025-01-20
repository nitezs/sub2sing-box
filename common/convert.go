package common

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/nitezs/sub2sing-box/constant"
	"github.com/nitezs/sub2sing-box/model"
	"github.com/nitezs/sub2sing-box/parser"
	"github.com/nitezs/sub2sing-box/util"
	box "github.com/sagernet/sing-box"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/include"
	"github.com/sagernet/sing-box/option"
	J "github.com/sagernet/sing/common/json"
)

var globalCtx = box.Context(context.Background(), include.InboundRegistry(), include.OutboundRegistry(), include.EndpointRegistry())

func Convert(
	subscriptions []string,
	proxies []string,
	templatePath string,
	delete string,
	rename map[string]string,
	enableGroup bool,
	groupType string,
	sortKey string,
	sortType string,
) (string, error) {
	result := ""
	var err error

	outbounds, err := ConvertSubscriptionsToSProxy(subscriptions)
	if err != nil {
		return "", err
	}
	for _, proxy := range proxies {
		p, err := ConvertCProxyToSProxy(proxy)
		if err != nil {
			return "", err
		}
		outbounds = append(outbounds, p)
	}

	if delete != "" {
		outbounds, err = DeleteProxy(outbounds, delete)
		if err != nil {
			return "", err
		}
	}

	for k, v := range rename {
		outbounds, err = RenameProxy(outbounds, k, v)
		if err != nil {
			return "", err
		}
	}

	set := make(map[string]bool)
	deduplicatedOutbounds := make([]model.Outbound, 0)
	for _, p := range outbounds {
		jsonBytes, err := json.Marshal(p)
		if err != nil {
			return "", err
		}
		if _, exists := set[string(jsonBytes)]; !exists {
			set[string(jsonBytes)] = true
			deduplicatedOutbounds = append(deduplicatedOutbounds, p)
		}
	}
	outbounds = deduplicatedOutbounds

	tagSet := make(map[string]bool)
	for i, p := range outbounds {
		if _, exists := tagSet[p.Tag]; exists {
			count := 1
			for {
				newTag := fmt.Sprintf("%s %d", p.Tag, count)
				if _, exists := tagSet[newTag]; !exists {
					outbounds[i].Tag = newTag
					break
				} else {
					count++
				}
			}
		}
	}

	if enableGroup {
		outbounds = AddCountryGroup(outbounds, groupType, sortKey, sortType)
	}
	if templatePath != "" {
		templateData, err := ReadTemplate(templatePath)
		if err != nil {
			return "", err
		}
		reg := regexp.MustCompile("\"<[A-Za-z]{2}>\"")
		group := false
		for _, v := range model.CountryEnglishName {
			if strings.Contains(templateData, v) {
				group = true
			}
		}
		if !enableGroup && (reg.MatchString(templateData) || strings.Contains(templateData, constant.AllCountryTags) || group) {
			outbounds = AddCountryGroup(outbounds, groupType, sortKey, sortType)
		}
		var template model.Options
		if template, err = J.UnmarshalExtendedContext[model.Options](globalCtx, []byte(templateData)); err != nil {
			return "", err
		}
		result, err = MergeTemplate(outbounds, &template)
		if err != nil {
			return "", err
		}
	} else {
		outboundJsons := make([]string, 0)
		for _, p := range outbounds {
			b, err := json.Marshal(p)
			if err != nil {
				return "", err
			}
			outboundJsons = append(outboundJsons, string(b))
		}
		result = fmt.Sprintf("[%s]", strings.Join(outboundJsons, ","))
		if err != nil {
			return "", err
		}
	}

	return string(result), nil
}

func AddCountryGroup(proxies []model.Outbound, groupType string, sortKey string, sortType string) []model.Outbound {
	newGroup := make(map[string]model.Outbound)
	for _, p := range proxies {
		if p.Type != C.TypeSelector && p.Type != C.TypeURLTest {
			country := model.GetContryName(p.Tag)
			if group, ok := newGroup[country]; ok {
				AppendOutbound(&group, p.Tag)
				newGroup[country] = group
			} else {
				if groupType == C.TypeSelector || groupType == "" {
					newGroup[country] = model.Outbound{
						Tag:  country,
						Type: groupType,
						Options: option.SelectorOutboundOptions{
							Outbounds:                 []string{p.Tag},
							InterruptExistConnections: true,
						},
					}
				} else if groupType == C.TypeURLTest {
					newGroup[country] = model.Outbound{
						Tag:  country,
						Type: groupType,
						Options: option.URLTestOutboundOptions{
							Outbounds:                 []string{p.Tag},
							InterruptExistConnections: true,
						},
					}
				}
			}
		}
	}
	var groups []model.Outbound
	for _, p := range newGroup {
		groups = append(groups, p)
	}
	if sortType == "asc" {
		switch sortKey {
		case "tag":
			sort.Sort(model.SortByTag(groups))
		case "num":
			sort.Sort(model.SortByNumber(groups))
		default:
			sort.Sort(model.SortByTag(groups))
		}
	} else {
		switch sortKey {
		case "tag":
			sort.Sort(sort.Reverse(model.SortByTag(groups)))
		case "num":
			sort.Sort(sort.Reverse(model.SortByNumber(groups)))
		default:
			sort.Sort(sort.Reverse(model.SortByTag(groups)))
		}
	}
	return append(proxies, groups...)
}

func ReadTemplate(template string) (string, error) {
	var data string
	var err error
	isNetworkFile, _ := regexp.MatchString(`^https?://`, template)
	if isNetworkFile {
		data, err = util.Fetch(template, 3)
		if err != nil {
			return "", err
		}
		return data, nil
	} else {
		if !strings.Contains(template, string(filepath.Separator)) {
			path := filepath.Join("templates", template)
			if _, err := os.Stat(path); err == nil {
				template = path
			}
		}
		dataBytes, err := os.ReadFile(template)
		if err != nil {
			return "", err
		}
		return string(dataBytes), nil
	}
}

func MergeTemplate(outbounds []model.Outbound, template *model.Options) (string, error) {
	var err error
	proxyTags := make([]string, 0)
	groupTags := make([]string, 0)
	groups := make(map[string]model.Outbound)
	for _, p := range outbounds {
		if model.IsCountryGroup(p.Tag) {
			groupTags = append(groupTags, p.Tag)
			reg := regexp.MustCompile("[A-Za-z]{2}")
			country := reg.FindString(p.Tag)
			groups[country] = p
		} else {
			proxyTags = append(proxyTags, p.Tag)
		}
	}
	reg := regexp.MustCompile("<[A-Za-z]{2}>")
	for i, outbound := range template.Outbounds {
		if outbound.Type == C.TypeSelector || outbound.Type == C.TypeURLTest {
			var parsedOutbound []string = make([]string, 0)
			for _, o := range GetOutbounds(&outbound) {
				if o == constant.AllProxyTags {
					parsedOutbound = append(parsedOutbound, proxyTags...)
				} else if o == constant.AllCountryTags {
					parsedOutbound = append(parsedOutbound, groupTags...)
				} else if reg.MatchString(o) {
					country := strings.ToUpper(strings.Trim(reg.FindString(o), "<>"))
					if group, ok := groups[country]; ok {
						parsedOutbound = append(parsedOutbound, GetOutbounds(&group)...)
					}
				} else {
					parsedOutbound = append(parsedOutbound, o)
				}
			}
			SetOutbounds(&template.Outbounds[i], parsedOutbound)
		}
	}
	template.Outbounds = append(template.Outbounds, outbounds...)

	for i := range template.DNS.Rules {
		if template.DNS.Rules[i].Type == "" {
			template.DNS.Rules[i].Type = C.RuleTypeDefault
		}
	}

	data, err := json.Marshal(template)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ConvertCProxyToSProxy(proxy string) (model.Outbound, error) {
	for prefix, parseFunc := range parser.ParserMap {
		if strings.HasPrefix(proxy, prefix) {
			proxy, err := parseFunc(proxy)
			if err != nil {
				return model.Outbound{}, err
			}
			return proxy, nil
		}
	}
	return model.Outbound{}, errors.New("unknown proxy format")
}

func ConvertSubscriptionsToSProxy(urls []string) ([]model.Outbound, error) {
	proxyList := make([]model.Outbound, 0)
	for _, url := range urls {
		data, err := util.Fetch(url, 3)
		if err != nil {
			return nil, err
		}
		proxy := data
		if !strings.Contains(data, "://") {
			proxy, err = util.DecodeBase64(data)
		}
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

func DeleteProxy(proxies []model.Outbound, regex string) ([]model.Outbound, error) {
	reg, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}
	var newProxies []model.Outbound
	for _, p := range proxies {
		if !reg.MatchString(p.Tag) {
			newProxies = append(newProxies, p)
		}
	}
	return newProxies, nil
}

func RenameProxy(proxies []model.Outbound, regex string, replaceText string) ([]model.Outbound, error) {
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

func SetOutbounds(outbound *model.Outbound, outbounds []string) {
	switch v := outbound.Options.(type) {
	case option.SelectorOutboundOptions:
		v.Outbounds = outbounds
		outbound.Options = v
	case option.URLTestOutboundOptions:
		v.Outbounds = outbounds
		outbound.Options = v
	}
}

func AppendOutbound(outbound *model.Outbound, outboundTag string) {
	switch v := outbound.Options.(type) {
	case option.SelectorOutboundOptions:
		v.Outbounds = append(v.Outbounds, outboundTag)
		outbound.Options = v
	case option.URLTestOutboundOptions:
		v.Outbounds = append(v.Outbounds, outboundTag)
		outbound.Options = v
	}
}

func GetOutbounds(outbound *model.Outbound) []string {
	switch v := outbound.Options.(type) {
	case option.SelectorOutboundOptions:
		return v.Outbounds
	case option.URLTestOutboundOptions:
		return v.Outbounds
	}
	return nil
}
