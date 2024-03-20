package parser

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
	"sub2sing-box/model"
)

func ParseTrojan(proxy string) (model.Outbound, error) {
	if !strings.HasPrefix(proxy, "trojan://") {
		return model.Outbound{}, errors.New("invalid trojan Url")
	}
	parts := strings.SplitN(strings.TrimPrefix(proxy, "trojan://"), "@", 2)
	if len(parts) != 2 {
		return model.Outbound{}, errors.New("invalid trojan Url")
	}
	serverInfo := strings.SplitN(parts[1], "#", 2)
	serverAndPortAndParams := strings.SplitN(serverInfo[0], "?", 2)
	serverAndPort := strings.SplitN(serverAndPortAndParams[0], ":", 2)
	params, err := url.ParseQuery(serverAndPortAndParams[1])
	if err != nil {
		return model.Outbound{}, err
	}
	if len(serverAndPort) != 2 {
		return model.Outbound{}, errors.New("invalid trojan Url")
	}
	port, err := strconv.Atoi(strings.TrimSpace(serverAndPort[1]))
	if err != nil {
		return model.Outbound{}, err
	}
	remarks := ""
	if len(serverInfo) == 2 {
		remarks, _ = url.QueryUnescape(strings.TrimSpace(serverInfo[1]))
	} else {
		remarks = serverAndPort[0]
	}
	server := strings.TrimSpace(serverAndPort[0])
	password := strings.TrimSpace(parts[0])
	result := model.Outbound{
		Type: "trojan",
		Tag:  remarks,
		TrojanOptions: model.TrojanOutboundOptions{
			ServerOptions: model.ServerOptions{
				Server:     server,
				ServerPort: uint16(port),
			},
			Password: password,
			Network:  params.Get("type"),
		},
	}
	if params.Get("security") == "xtls" || params.Get("security") == "tls" {
		var alpn []string
		if strings.Contains(params.Get("alpn"), ",") {
			alpn = strings.Split(params.Get("alpn"), ",")
		} else {
			alpn = nil
		}
		result.TrojanOptions.OutboundTLSOptionsContainer = model.OutboundTLSOptionsContainer{
			TLS: &model.OutboundTLSOptions{
				Enabled:    true,
				ALPN:       alpn,
				ServerName: params.Get("sni"),
			},
		}
	}
	if params.Get("security") == "reality" {
		result.TrojanOptions.OutboundTLSOptionsContainer = model.OutboundTLSOptionsContainer{
			TLS: &model.OutboundTLSOptions{
				Enabled:    true,
				ServerName: params.Get("sni"),
				Reality: &model.OutboundRealityOptions{
					Enabled:   true,
					PublicKey: params.Get("pbk"),
					ShortID:   params.Get("sid"),
				},
				UTLS: &model.OutboundUTLSOptions{
					Enabled:     params.Get("fp") != "",
					Fingerprint: params.Get("fp"),
				},
			},
		}
	}
	if params.Get("type") == "ws" {
		result.TrojanOptions.Transport = &model.V2RayTransportOptions{
			Type: "ws",
			WebsocketOptions: model.V2RayWebsocketOptions{
				Path: params.Get("path"),
				Headers: map[string]string{
					"Host": params.Get("host"),
				},
			},
		}
	}
	if params.Get("type") == "http" {
		result.TrojanOptions.Transport = &model.V2RayTransportOptions{
			Type: "http",
			HTTPOptions: model.V2RayHTTPOptions{
				Host: []string{params.Get("host")},
				Path: params.Get("path"),
			},
		}
	}
	if params.Get("type") == "quic" {
		result.TrojanOptions.Transport = &model.V2RayTransportOptions{
			Type:        "quic",
			QUICOptions: model.V2RayQUICOptions{},
		}
	}
	if params.Get("type") == "grpc" {
		result.TrojanOptions.Transport = &model.V2RayTransportOptions{
			Type: "grpc",
			GRPCOptions: model.V2RayGRPCOptions{
				ServiceName: params.Get("serviceName"),
			},
		}
	}
	return result, nil
}
