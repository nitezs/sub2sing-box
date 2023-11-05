package common

import (
	"sub2sing-box/model"

	"gopkg.in/yaml.v3"
)

func ParseClashRules(singBox *model.SingBox, config []byte) error {
	return nil
}

func ParseClashProxies(singBox *model.SingBox, config []byte) error {
	// parse config
	var clashConfig model.Clash
	yaml.Unmarshal(config, &clashConfig)

	// parse proxies
	for _, proxy := range clashConfig.Proxies {
		var newProxy *model.Outbound
		switch proxy.Type {
		case "ss":
			newProxy = ParseClashShadowsocks(proxy)
		}

		if newProxy != nil {
			// add to sing box config
			singBox.Outbounds = append(singBox.Outbounds, *newProxy)
		}
	}

	return nil
}
