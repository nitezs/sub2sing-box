package parser

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sub2sing-box/internal/model"
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
		Tag:  remarks,
		Trojan: model.Trojan{
			Server:     server,
			ServerPort: uint16(port),
			Password:   password,
			Network:    params.Get("type"),
		},
	}
	if params.Get("security") == "xtls" || params.Get("security") == "tls" {
		result.Trojan.TLS = &model.OutboundTLSOptions{
			Enabled:    true,
			ALPN:       strings.Split(params.Get("alpn"), ","),
			ServerName: params.Get("sni"),
		}
	}
	if params.Get("security") == "reality" {
		result.Trojan.TLS = &model.OutboundTLSOptions{
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
		}
	}
	if params.Get("type") == "ws" {
		result.Trojan.Transport = &model.V2RayTransportOptions{
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
		result.Trojan.Transport = &model.V2RayTransportOptions{
			Type: "http",
			HTTPOptions: model.V2RayHTTPOptions{
				Host: []string{params.Get("host")},
				Path: params.Get("path"),
			},
		}
	}
	if params.Get("type") == "quic" {
		result.Trojan.Transport = &model.V2RayTransportOptions{
			Type:        "quic",
			QUICOptions: model.V2RayQUICOptions{},
		}
	}
	if params.Get("type") == "grpc" {
		result.Trojan.Transport = &model.V2RayTransportOptions{
			Type: "grpc",
			GRPCOptions: model.V2RayGRPCOptions{
				ServiceName: params.Get("serviceName"),
			},
		}
	}
	return result, nil
}
