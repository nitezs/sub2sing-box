package common

import "sub2sing-box/model"

func AddAllProxiesSelector(singBox *model.SingBox) {
	var selector model.Outbound
	selector.Tag = "proxies"
	selector.Type = "selector"
	selector.Outbounds = []string{}
	for _, outbound := range singBox.Outbounds {
		selector.Outbounds = append(selector.Outbounds, outbound.Tag)
	}
	singBox.Outbounds = append(singBox.Outbounds, selector)
}
