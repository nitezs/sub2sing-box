package parser

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/nitezs/sub2sing-box/constant"
	"github.com/nitezs/sub2sing-box/model"
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

	if security == "xtls" || security == "tls" || sni != "" {
		result.TrojanOptions.OutboundTLSOptionsContainer = model.OutboundTLSOptionsContainer{
			TLS: &model.OutboundTLSOptions{
				Enabled:    true,
				ALPN:       alpn,
				ServerName: sni,
				Insecure:   allowInsecure == "1",
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
				Insecure: allowInsecure == "1",
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
