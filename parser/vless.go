package parser

import (
	"net/url"
	"strings"
	"sub2sing-box/constant"
	"sub2sing-box/model"
)

func ParseVless(proxy string) (model.Outbound, error) {
	if !strings.HasPrefix(proxy, constant.VLESSPrefix) {
		return model.Outbound{}, &ParseError{Type: ErrInvalidPrefix, Raw: proxy}
	}

	urlParts := strings.SplitN(strings.TrimPrefix(proxy, constant.VLESSPrefix), "@", 2)
	if len(urlParts) != 2 {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing character '@' in url",
			Raw:     proxy,
		}
	}

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
	port, err := ParsePort(portStr)
	if err != nil {
		return model.Outbound{}, err
	}

	params, err := url.ParseQuery(serverAndPortAndParams[1])
	if err != nil {
		return model.Outbound{}, &ParseError{
			Type:    ErrCannotParseParams,
			Raw:     proxy,
			Message: err.Error(),
		}
	}

	remarks := ""
	if len(serverInfo) == 2 {
		if strings.Contains(serverInfo[1], "|") {
			remarks = strings.SplitN(serverInfo[1], "|", 2)[1]
		} else {
			remarks, err = url.QueryUnescape(serverInfo[1])
			if err != nil {
				return model.Outbound{}, &ParseError{
					Type:    ErrCannotParseParams,
					Raw:     proxy,
					Message: err.Error(),
				}
			}
		}
	} else {
		remarks, err = url.QueryUnescape(server)
		if err != nil {
			return model.Outbound{}, err
		}
	}

	uuid := strings.TrimSpace(urlParts[0])
	flow, security, alpnStr, sni, insecure, fp, pbk, sid, path, host, serviceName := params.Get("flow"), params.Get("security"), params.Get("alpn"), params.Get("sni"), params.Get("allowInsecure"), params.Get("fp"), params.Get("pbk"), params.Get("sid"), params.Get("path"), params.Get("host"), params.Get("serviceName")

	enableUTLS := fp != ""
	insecureBool := insecure == "1"
	var alpn []string
	if strings.Contains(alpnStr, ",") {
		alpn = strings.Split(alpnStr, ",")
	} else {
		alpn = nil
	}

	result := model.Outbound{
		Type: "vless",
		Tag:  remarks,
		VLESSOptions: model.VLESSOutboundOptions{
			ServerOptions: model.ServerOptions{
				Server:     server,
				ServerPort: port,
			},
			UUID: uuid,
			Flow: flow,
		},
	}

	if security == "tls" {
		result.VLESSOptions.OutboundTLSOptionsContainer = model.OutboundTLSOptionsContainer{
			TLS: &model.OutboundTLSOptions{
				Enabled:    true,
				ALPN:       alpn,
				ServerName: sni,
				Insecure:   insecureBool,
			},
		}
		result.VLESSOptions.OutboundTLSOptionsContainer.TLS.UTLS = &model.OutboundUTLSOptions{
			Enabled:     enableUTLS,
			Fingerprint: fp,
		}
	}

	if security == "reality" {
		result.VLESSOptions.OutboundTLSOptionsContainer.TLS.Reality = &model.OutboundRealityOptions{
			Enabled:   true,
			PublicKey: pbk,
			ShortID:   sid,
		}
	}

	if params.Get("type") == "ws" {
		result.VLESSOptions.Transport = &model.V2RayTransportOptions{
			Type: "ws",
			WebsocketOptions: model.V2RayWebsocketOptions{
				Path: path,
			},
		}
		if host != "" {
			if result.VLESSOptions.Transport.WebsocketOptions.Headers == nil {
				result.VLESSOptions.Transport.WebsocketOptions.Headers = make(map[string]string)
			}
			result.VLESSOptions.Transport.WebsocketOptions.Headers["Host"] = host
		}
	}

	if params.Get("type") == "quic" {
		result.VLESSOptions.Transport = &model.V2RayTransportOptions{
			Type:        "quic",
			QUICOptions: model.V2RayQUICOptions{},
		}
	}

	if params.Get("type") == "grpc" {
		result.VLESSOptions.Transport = &model.V2RayTransportOptions{
			Type: "grpc",
			GRPCOptions: model.V2RayGRPCOptions{
				ServiceName: serviceName,
			},
		}
	}

	if params.Get("type") == "http" {
		hosts, err := url.QueryUnescape(host)
		if err != nil {
			return model.Outbound{}, &ParseError{
				Type:    ErrCannotParseParams,
				Raw:     proxy,
				Message: err.Error(),
			}
		}
		result.VLESSOptions.Transport = &model.V2RayTransportOptions{
			Type: "http",
			HTTPOptions: model.V2RayHTTPOptions{
				Host: strings.Split(hosts, ","),
			},
		}
	}
	return result, nil
}
