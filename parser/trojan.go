package parser

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/nitezs/sub2sing-box/constant"
	"github.com/nitezs/sub2sing-box/model"
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common/json/badoption"
)

func ParseTrojan(proxy string) (model.Outbound, error) {
	if !strings.HasPrefix(proxy, constant.TrojanPrefix) {
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

	password := link.User.Username()
	server := link.Hostname()
	if server == "" {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing server host",
			Raw:     proxy,
		}
	}
	portStr := link.Port()
	if portStr == "" {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing server port",
			Raw:     proxy,
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

	remarks := link.Fragment
	if remarks == "" {
		remarks = fmt.Sprintf("%s:%s", server, portStr)
	}
	remarks = strings.TrimSpace(remarks)

	query := link.Query()
	network, security, alpnStr, sni, pbk, sid, fp, path, host, serviceName, allowInsecure := query.Get("type"), query.Get("security"), query.Get("alpn"), query.Get("sni"), query.Get("pbk"), query.Get("sid"), query.Get("fp"), query.Get("path"), query.Get("host"), query.Get("serviceName"), query.Get("allowInsecure")

	var alpn []string
	if strings.Contains(alpnStr, ",") {
		alpn = strings.Split(alpnStr, ",")
	} else {
		alpn = nil
	}

	enableUTLS := fp != ""

	outboundOptions := option.TrojanOutboundOptions{
		ServerOptions: option.ServerOptions{
			Server:     server,
			ServerPort: port,
		},
		Password: password,
		Network:  option.NetworkList(network),
	}

	if security == "xtls" || security == "tls" || sni != "" {
		outboundOptions.OutboundTLSOptionsContainer = option.OutboundTLSOptionsContainer{
			TLS: &option.OutboundTLSOptions{
				Enabled:    true,
				ALPN:       alpn,
				ServerName: sni,
				Insecure:   allowInsecure == "1",
			},
		}
	}

	if security == "reality" {
		outboundOptions.OutboundTLSOptionsContainer = option.OutboundTLSOptionsContainer{
			TLS: &option.OutboundTLSOptions{
				Enabled:    true,
				ServerName: sni,
				Reality: &option.OutboundRealityOptions{
					Enabled:   true,
					PublicKey: pbk,
					ShortID:   sid,
				},
				UTLS: &option.OutboundUTLSOptions{
					Enabled:     enableUTLS,
					Fingerprint: fp,
				},
				Insecure: allowInsecure == "1",
			},
		}
	}

	if network == "ws" {
		outboundOptions.Transport = &option.V2RayTransportOptions{
			Type: "ws",
			WebsocketOptions: option.V2RayWebsocketOptions{
				Path: path,
				Headers: badoption.HTTPHeader{
					"Host": {host},
				},
			},
		}
	}

	if network == "http" {
		outboundOptions.Transport = &option.V2RayTransportOptions{
			Type: "http",
			HTTPOptions: option.V2RayHTTPOptions{
				Host: []string{host},
				Path: path,
			},
		}
	}

	if network == "quic" {
		outboundOptions.Transport = &option.V2RayTransportOptions{
			Type:        "quic",
			QUICOptions: option.V2RayQUICOptions{},
		}
	}

	if network == "grpc" {
		outboundOptions.Transport = &option.V2RayTransportOptions{
			Type: "grpc",
			GRPCOptions: option.V2RayGRPCOptions{
				ServiceName: serviceName,
			},
		}
	}

	result := model.Outbound{Outbound: option.Outbound{
		Type:    "trojan",
		Tag:     remarks,
		Options: outboundOptions,
	}}

	return result, nil
}
