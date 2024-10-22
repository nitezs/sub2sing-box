package parser

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/nitezs/sub2sing-box/constant"
	"github.com/nitezs/sub2sing-box/model"
	"github.com/sagernet/sing-box/option"
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

	result := model.Outbound{Outbound: option.Outbound{
		Type: "vless",
		Tag:  remarks,
		VLESSOptions: option.VLESSOutboundOptions{
			ServerOptions: option.ServerOptions{
				Server:     server,
				ServerPort: port,
			},
			UUID: uuid,
			Flow: flow,
		},
	}}

	if security == "tls" {
		result.VLESSOptions.OutboundTLSOptionsContainer = option.OutboundTLSOptionsContainer{
			TLS: &option.OutboundTLSOptions{
				Enabled:    true,
				ALPN:       alpn,
				ServerName: sni,
				Insecure:   insecureBool,
			},
		}
	}

	if security == "reality" {
		result.VLESSOptions.OutboundTLSOptionsContainer = option.OutboundTLSOptionsContainer{
			TLS: &option.OutboundTLSOptions{
				Enabled:    true,
				ALPN:       alpn,
				ServerName: sni,
				Insecure:   insecureBool,
				Reality: &option.OutboundRealityOptions{
					Enabled:   true,
					PublicKey: pbk,
					ShortID:   sid,
				},
				UTLS: &option.OutboundUTLSOptions{
					Enabled:     true,
					Fingerprint: fp,
				},
			},
		}
	}

	if _type == "ws" {
		result.VLESSOptions.Transport = &option.V2RayTransportOptions{
			Type: "ws",
			WebsocketOptions: option.V2RayWebsocketOptions{
				Path: path,
			},
		}
		if host != "" {
			if result.VLESSOptions.Transport.WebsocketOptions.Headers == nil {
				result.VLESSOptions.Transport.WebsocketOptions.Headers = make(map[string]option.Listable[string])
			}
			result.VLESSOptions.Transport.WebsocketOptions.Headers["Host"] = option.Listable[string]{host}
		}
	}

	if _type == "quic" {
		result.VLESSOptions.Transport = &option.V2RayTransportOptions{
			Type:        "quic",
			QUICOptions: option.V2RayQUICOptions{},
		}
	}

	if _type == "grpc" {
		result.VLESSOptions.Transport = &option.V2RayTransportOptions{
			Type: "grpc",
			GRPCOptions: option.V2RayGRPCOptions{
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
		result.VLESSOptions.Transport = &option.V2RayTransportOptions{
			Type: "http",
			HTTPOptions: option.V2RayHTTPOptions{
				Host: strings.Split(hosts, ","),
			},
		}
	}

	if enableUTLS {
		result.VLESSOptions.OutboundTLSOptionsContainer.TLS.UTLS = &option.OutboundUTLSOptions{
			Enabled:     enableUTLS,
			Fingerprint: fp,
		}
	}

	return result, nil
}
