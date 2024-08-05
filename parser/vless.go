package parser

import (
	"fmt"
	"net/url"
	"strings"
	"sub2sing-box/constant"
	"sub2sing-box/model"
)

func ParseVless(proxy string) (model.Outbound, error) {
	if !strings.HasPrefix(proxy, constant.VLESSPrefix) {
		return model.Outbound{}, &ParseError{Type: ErrInvalidPrefix, Raw: proxy}
	}

	link, err := url.Parse(proxy)
	if err != nil {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "url parse error",
			Raw:     proxy,
		}
	}

	server := link.Hostname()
	if server == "" {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing server host",
			Raw:     proxy,
		}
	}
	portStr := link.Port()
	port, err := ParsePort(portStr)
	if err != nil {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidPort,
			Message: err.Error(),
			Raw:     proxy,
		}
	}

	query := link.Query()
	uuid := link.User.Username()
	flow, security, alpnStr, sni, insecure, fp, pbk, sid, path, host, serviceName, _type := query.Get("flow"), query.Get("security"), query.Get("alpn"), query.Get("sni"), query.Get("allowInsecure"), query.Get("fp"), query.Get("pbk"), query.Get("sid"), query.Get("path"), query.Get("host"), query.Get("serviceName"), query.Get("type")

	enableUTLS := fp != ""
	insecureBool := insecure == "1"
	var alpn []string
	if strings.Contains(alpnStr, ",") {
		alpn = strings.Split(alpnStr, ",")
	} else {
		alpn = nil
	}
	remarks := link.Fragment
	if remarks == "" {
		remarks = fmt.Sprintf("%s:%s", server, portStr)
	}
	remarks = strings.TrimSpace(remarks)

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
	}

	if security == "reality" {
		result.VLESSOptions.OutboundTLSOptionsContainer = model.OutboundTLSOptionsContainer{
			TLS: &model.OutboundTLSOptions{
				Enabled:    true,
				ALPN:       alpn,
				ServerName: sni,
				Insecure:   insecureBool,
				Reality: &model.OutboundRealityOptions{
					Enabled:   true,
					PublicKey: pbk,
					ShortID:   sid,
				},
			},
		}
	}

	if _type == "ws" {
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

	if _type == "quic" {
		result.VLESSOptions.Transport = &model.V2RayTransportOptions{
			Type:        "quic",
			QUICOptions: model.V2RayQUICOptions{},
		}
	}

	if _type == "grpc" {
		result.VLESSOptions.Transport = &model.V2RayTransportOptions{
			Type: "grpc",
			GRPCOptions: model.V2RayGRPCOptions{
				ServiceName: serviceName,
			},
		}
	}

	if _type == "http" {
		hosts, err := url.QueryUnescape(host)
		if err != nil {
			return model.Outbound{}, &ParseError{
				Type:    ErrInvalidStruct,
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

	if enableUTLS {
		result.VLESSOptions.OutboundTLSOptionsContainer.TLS.UTLS = &model.OutboundUTLSOptions{
			Enabled:     enableUTLS,
			Fingerprint: fp,
		}
	}

	return result, nil
}
