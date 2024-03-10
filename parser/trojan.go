package parser

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sub2sing-box/model"
)

func ParseTrojan(proxy string) (model.Proxy, error) {
	if !strings.HasPrefix(proxy, "trojan://") {
		return model.Proxy{}, fmt.Errorf("invalid trojan Url")
	}
	parts := strings.SplitN(strings.TrimPrefix(proxy, "trojan://"), "@", 2)
	if len(parts) != 2 {
		return model.Proxy{}, fmt.Errorf("invalid trojan Url")
	}
	serverInfo := strings.SplitN(parts[1], "#", 2)
	serverAndPortAndParams := strings.SplitN(serverInfo[0], "?", 2)
	serverAndPort := strings.SplitN(serverAndPortAndParams[0], ":", 2)
	params, err := url.ParseQuery(serverAndPortAndParams[1])
	if err != nil {
		return model.Proxy{}, err
	}
	if len(serverAndPort) != 2 {
		return model.Proxy{}, fmt.Errorf("invalid trojan")
	}
	port, err := strconv.Atoi(strings.TrimSpace(serverAndPort[1]))
	if err != nil {
		return model.Proxy{}, err
	}
	remarks := ""
	if len(serverInfo) == 2 {
		remarks, _ = url.QueryUnescape(strings.TrimSpace(serverInfo[1]))
	} else {
		remarks = serverAndPort[0]
	}
	server := strings.TrimSpace(serverAndPort[0])
	password := strings.TrimSpace(parts[0])
	result := model.Proxy{
		Type: "trojan",
		Trojan: model.Trojan{
			Tag:        remarks,
			Server:     server,
			ServerPort: uint16(port),
			TLS: &model.OutboundTLSOptions{
				Enabled:    true,
				ServerName: params.Get("sni"),
			},
			Password: password,
			Network:  "tcp",
		},
	}
	return result, nil
}
