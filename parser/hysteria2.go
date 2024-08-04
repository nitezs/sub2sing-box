package parser

import (
	"fmt"
	"net/url"
	"strings"
	"sub2sing-box/constant"
	"sub2sing-box/model"
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
	network, obfs, obfsPassword, pinSHA256, insecure, sni := query.Get("network"), query.Get("obfs"), query.Get("obfs-password"), query.Get("pinSHA256"), query.Get("insecure"), query.Get("sni")
	enableTLS := pinSHA256 != ""
	insecureBool := insecure == "1"
	remarks := link.Fragment
	if remarks == "" {
		remarks = fmt.Sprintf("%s:%s", server, portStr)
	}
	remarks = strings.TrimSpace(remarks)

	result := model.Outbound{
		Type: "hysteria2",
		Tag:  strings.TrimSpace(remarks),
		Hysteria2Options: model.Hysteria2OutboundOptions{
			ServerOptions: model.ServerOptions{
				Server:     server,
				ServerPort: port,
			},
			Password: password,
			Obfs: &model.Hysteria2Obfs{
				Type:     obfs,
				Password: obfsPassword,
			},
			OutboundTLSOptionsContainer: model.OutboundTLSOptionsContainer{
				TLS: &model.OutboundTLSOptions{Enabled: enableTLS,
					Insecure:    insecureBool,
					ServerName:  sni,
					Certificate: []string{pinSHA256}},
			},
			Network: network,
		},
	}
	return result, nil
}
