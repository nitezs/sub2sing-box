package common

import "sub2sing-box/model"

func ParseClashShadowsocks(proxy model.Proxy) *model.Outbound {
	network := "tcp"
	if proxy.UDP {
		network = "udp"
	}
	return &model.Outbound{
		Type:       "shadowsocks",
		Tag:        proxy.Name,
		Server:     proxy.Server,
		ServerPort: proxy.Port,
		Method:     proxy.Cipher,
		Password:   proxy.Password,
		Plugin:     proxy.Plugin,
		PluginOpts: proxy.PluginOpts,
		Network:    network,
		UDPOverTCP: proxy.UDPOverTCP,
	}
}
