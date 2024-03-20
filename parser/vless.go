package parser

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
	"sub2sing-box/model"
)

func ParseVless(proxy string) (model.Outbound, error) {
	if !strings.HasPrefix(proxy, "vless://") {
		return model.Outbound{}, errors.New("invalid vless Url")
	}
	parts := strings.SplitN(strings.TrimPrefix(proxy, "vless://"), "@", 2)
	if len(parts) != 2 {
		return model.Outbound{}, errors.New("invalid vless Url")
	}
	serverInfo := strings.SplitN(parts[1], "#", 2)
	serverAndPortAndParams := strings.SplitN(serverInfo[0], "?", 2)
	serverAndPort := strings.SplitN(serverAndPortAndParams[0], ":", 2)
	params, err := url.ParseQuery(serverAndPortAndParams[1])
	if err != nil {
		return model.Outbound{}, err
	}
	if len(serverAndPort) != 2 {
		return model.Outbound{}, errors.New("invalid vless Url")
	}
	port, err := strconv.Atoi(strings.TrimSpace(serverAndPort[1]))
	if err != nil {
		return model.Outbound{}, err
	}
	remarks := ""
	if len(serverInfo) == 2 {
		if strings.Contains(serverInfo[1], "|") {
			remarks = strings.SplitN(serverInfo[1], "|", 2)[1]
		} else {
			remarks, err = url.QueryUnescape(serverInfo[1])
			if err != nil {
				return model.Outbound{}, err
			}
		}
	} else {
		remarks, err = url.QueryUnescape(serverAndPort[0])
		if err != nil {
			return model.Outbound{}, err
		}
	}
	server := strings.TrimSpace(serverAndPort[0])
	uuid := strings.TrimSpace(parts[0])
	result := model.Outbound{
		Type: "vless",
		Tag:  remarks,
		VLESSOptions: model.VLESSOutboundOptions{
			ServerOptions: model.ServerOptions{
				Server:     server,
				ServerPort: uint16(port),
			},
			UUID: uuid,
			Flow: params.Get("flow"),
		},
	}
	if params.Get("security") == "tls" {
		var alpn []string
		if strings.Contains(params.Get("alpn"), ",") {
			alpn = strings.Split(params.Get("alpn"), ",")
		} else {
			alpn = nil
		}
		result.VLESSOptions.OutboundTLSOptionsContainer = model.OutboundTLSOptionsContainer{
			TLS: &model.OutboundTLSOptions{
				Enabled:    true,
				ALPN:       alpn,
				ServerName: params.Get("sni"),
				Insecure:   params.Get("allowInsecure") == "1",
			},
		}
		if params.Get("fp") != "" {
			result.VLESSOptions.OutboundTLSOptionsContainer.TLS.UTLS = &model.OutboundUTLSOptions{
				Enabled:     true,
				Fingerprint: params.Get("fp"),
			}
		}
	}
	if params.Get("security") == "reality" {
		result.VLESSOptions.OutboundTLSOptionsContainer.TLS.Reality = &model.OutboundRealityOptions{
			Enabled:   true,
			PublicKey: params.Get("pbk"),
			ShortID:   params.Get("sid"),
		}
	}
	if params.Get("type") == "ws" {
		result.VLESSOptions.Transport = &model.V2RayTransportOptions{
			Type: "ws",
			WebsocketOptions: model.V2RayWebsocketOptions{
				Path: params.Get("path"),
			},
		}
		if params.Get("host") != "" {
			result.VLESSOptions.Transport.WebsocketOptions.Headers["Host"] = params.Get("host")
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
				ServiceName: params.Get("serviceName"),
			},
		}
	}
	if params.Get("type") == "http" {
		host, err := url.QueryUnescape(params.Get("host"))
		if err != nil {
			return model.Outbound{}, err
		}
		result.VLESSOptions.Transport = &model.V2RayTransportOptions{
			Type: "http",
			HTTPOptions: model.V2RayHTTPOptions{
				Host: strings.Split(host, ","),
			},
		}
	}
	return result, nil
}
