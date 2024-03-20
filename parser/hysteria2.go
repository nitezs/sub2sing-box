package parser

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
	"sub2sing-box/model"
)

// hysteria2://letmein@example.com/?insecure=1&obfs=salamander&obfs-password=gawrgura&pinSHA256=deadbeef&sni=real.example.com

func ParseHysteria2(proxy string) (model.Outbound, error) {
	if !strings.HasPrefix(proxy, "hysteria2://") && !strings.HasPrefix(proxy, "hy2://") {
		return model.Outbound{}, errors.New("invalid hysteria2 Url")
	}
	parts := strings.SplitN(strings.TrimPrefix(proxy, "hysteria2://"), "@", 2)
	serverInfo := strings.SplitN(parts[1], "/?", 2)
	serverAndPort := strings.SplitN(serverInfo[0], ":", 2)
	if len(serverAndPort) == 1 {
		serverAndPort = append(serverAndPort, "443")
	} else if len(serverAndPort) != 2 {
		return model.Outbound{}, errors.New("invalid hysteria2 Url")
	}
	params, err := url.ParseQuery(serverInfo[1])
	if err != nil {
		return model.Outbound{}, errors.New("invalid hysteria2 Url")
	}
	port, err := strconv.Atoi(serverAndPort[1])
	if err != nil {
		return model.Outbound{}, errors.New("invalid hysteria2 Url")
	}
	remarks := params.Get("name")
	server := serverAndPort[0]
	password := parts[0]
	network := params.Get("network")
	result := model.Outbound{
		Type: "hysteria2",
		Tag:  remarks,
		Hysteria2Options: model.Hysteria2OutboundOptions{
			ServerOptions: model.ServerOptions{
				Server:     server,
				ServerPort: uint16(port),
			},
			Password: password,
			Obfs: &model.Hysteria2Obfs{
				Type:     params.Get("obfs"),
				Password: params.Get("obfs-password"),
			},
			OutboundTLSOptionsContainer: model.OutboundTLSOptionsContainer{
				TLS: &model.OutboundTLSOptions{Enabled: params.Get("pinSHA256") != "",
					Insecure:    params.Get("insecure") == "1",
					ServerName:  params.Get("sni"),
					Certificate: []string{params.Get("pinSHA256")}},
			},
			Network: network,
		},
	}
	return result, nil
}
