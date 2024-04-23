package parser

import (
	"net/url"
	"strings"
	"sub2sing-box/constant"
	"sub2sing-box/model"
)

func ParseTrojan(proxy string) (model.Outbound, error) {
	if !strings.HasPrefix(proxy, constant.TrojanPrefix) {
		return model.Outbound{}, &ParseError{Type: ErrInvalidPrefix, Raw: proxy}
	}

	proxy = strings.TrimPrefix(proxy, constant.TrojanPrefix)
	urlParts := strings.SplitN(proxy, "@", 2)
	if len(urlParts) != 2 {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing character '@' in url",
			Raw:     proxy,
		}
	}
	password := strings.TrimSpace(urlParts[0])

	serverInfo := strings.SplitN(urlParts[1], "#", 2)
	serverAndPortAndParams := strings.SplitN(serverInfo[0], "?", 2)
	if len(serverAndPortAndParams) != 2 {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing character '?' in url",
			Raw:     proxy,
		}
	}

	serverAndPort := strings.SplitN(serverAndPortAndParams[0], ":", 2)
	if len(serverAndPort) != 2 {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing server host or port",
			Raw:     proxy,
		}
	}
	server, portStr := serverAndPort[0], serverAndPort[1]

	params, err := url.ParseQuery(serverAndPortAndParams[1])
	if err != nil {
		return model.Outbound{}, &ParseError{
			Type:    ErrCannotParseParams,
			Raw:     proxy,
			Message: err.Error(),
		}
	}

	port, err := ParsePort(portStr)
	if err != nil {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidPort,
			Message: err.Error(),
			Raw:     proxy,
		}
	}

	remarks := ""
	if len(serverInfo) == 2 {
		remarks, _ = url.QueryUnescape(strings.TrimSpace(serverInfo[1]))
	} else {
		remarks = serverAndPort[0]
	}

	network, security, alpnStr, sni, pbk, sid, fp, path, host, serviceName := params.Get("type"), params.Get("security"), params.Get("alpn"), params.Get("sni"), params.Get("pbk"), params.Get("sid"), params.Get("fp"), params.Get("path"), params.Get("host"), params.Get("serviceName")

	var alpn []string
	if strings.Contains(alpnStr, ",") {
		alpn = strings.Split(alpnStr, ",")
	} else {
		alpn = nil
	}

	enableUTLS := fp != ""

	result := model.Outbound{
		Type: "trojan",
		Tag:  remarks,
		TrojanOptions: model.TrojanOutboundOptions{
			ServerOptions: model.ServerOptions{
				Server:     server,
				ServerPort: port,
			},
			Password: password,
			Network:  network,
		},
	}

	if security == "xtls" || security == "tls" {
		result.TrojanOptions.OutboundTLSOptionsContainer = model.OutboundTLSOptionsContainer{
			TLS: &model.OutboundTLSOptions{
				Enabled:    true,
				ALPN:       alpn,
				ServerName: sni,
			},
		}
	}

	if security == "reality" {
		result.TrojanOptions.OutboundTLSOptionsContainer = model.OutboundTLSOptionsContainer{
			TLS: &model.OutboundTLSOptions{
				Enabled:    true,
				ServerName: sni,
				Reality: &model.OutboundRealityOptions{
					Enabled:   true,
					PublicKey: pbk,
					ShortID:   sid,
				},
				UTLS: &model.OutboundUTLSOptions{
					Enabled:     enableUTLS,
					Fingerprint: fp,
				},
			},
		}
	}

	if network == "ws" {
		result.TrojanOptions.Transport = &model.V2RayTransportOptions{
			Type: "ws",
			WebsocketOptions: model.V2RayWebsocketOptions{
				Path: path,
				Headers: map[string]string{
					"Host": host,
				},
			},
		}
	}

	if network == "http" {
		result.TrojanOptions.Transport = &model.V2RayTransportOptions{
			Type: "http",
			HTTPOptions: model.V2RayHTTPOptions{
				Host: []string{host},
				Path: path,
			},
		}
	}

	if network == "quic" {
		result.TrojanOptions.Transport = &model.V2RayTransportOptions{
			Type:        "quic",
			QUICOptions: model.V2RayQUICOptions{},
		}
	}

	if network == "grpc" {
		result.TrojanOptions.Transport = &model.V2RayTransportOptions{
			Type: "grpc",
			GRPCOptions: model.V2RayGRPCOptions{
				ServiceName: serviceName,
			},
		}
	}
	return result, nil
}
