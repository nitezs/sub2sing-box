package parser

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/nitezs/sub2sing-box/constant"
	"github.com/nitezs/sub2sing-box/model"
	"github.com/sagernet/sing-box/option"
)

func ParseHysteria2(proxy string) (model.Outbound, error) {
	if !strings.HasPrefix(proxy, constant.Hysteria2Prefix1) &&
		!strings.HasPrefix(proxy, constant.Hysteria2Prefix2) {
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

	username := link.User.Username()
	password, exist := link.User.Password()
	if !exist {
		password = username
	}

	query := link.Query()
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
			Type: ErrInvalidPort,
			Raw:  portStr,
		}
	}
	network, obfs, obfsPassword, pinSHA256, insecure, sni, alpnStr := query.Get("network"), query.Get("obfs"), query.Get("obfs-password"), query.Get("pinSHA256"), query.Get("insecure"), query.Get("sni"), query.Get("alpn")
	insecureBool := insecure == "1"
	enableTLS := pinSHA256 != "" || sni != "" || alpnStr != ""

	var alpn []string
	alpnStr = strings.TrimSpace(alpnStr)
	if alpnStr != "" {
		alpn = strings.Split(alpnStr, ",")
	}

	remarks := link.Fragment
	if remarks == "" {
		remarks = fmt.Sprintf("%s:%s", server, portStr)
	}
	remarks = strings.TrimSpace(remarks)

	result := model.Outbound{
		Outbound: option.Outbound{
			Type: "hysteria2",
			Tag:  strings.TrimSpace(remarks),
			Hysteria2Options: option.Hysteria2OutboundOptions{
				ServerOptions: option.ServerOptions{
					Server:     server,
					ServerPort: port,
				},
				Password: password,
				OutboundTLSOptionsContainer: option.OutboundTLSOptionsContainer{
					TLS: &option.OutboundTLSOptions{
						Enabled:    enableTLS,
						Insecure:   insecureBool,
						ServerName: sni,
						ALPN:       alpn,
					},
				},
				Network: option.NetworkList(network),
			},
		},
	}
	if pinSHA256 != "" {
		result.Hysteria2Options.OutboundTLSOptionsContainer.TLS.Certificate = []string{pinSHA256}
	}
	if obfs != "" {
		result.Hysteria2Options.Obfs = &option.Hysteria2Obfs{
			Type:     obfs,
			Password: obfsPassword,
		}
	}
	return result, nil
}
