package parser

import (
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

	proxy = strings.TrimPrefix(proxy, constant.Hysteria2Prefix1)
	proxy = strings.TrimPrefix(proxy, constant.Hysteria2Prefix2)
	urlParts := strings.SplitN(proxy, "@", 2)
	if len(urlParts) != 2 {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing character '@' in url",
			Raw:     proxy,
		}
	}
	password := urlParts[0]

	serverInfo := strings.SplitN(urlParts[1], "/?", 2)
	if len(serverInfo) != 2 {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing params in url",
			Raw:     proxy,
		}
	}
	paramStr := serverInfo[1]

	serverAndPort := strings.SplitN(serverInfo[0], ":", 2)
	var server string
	var portStr string
	if len(serverAndPort) == 1 {
		portStr = "443"
	} else if len(serverAndPort) == 2 {
		server, portStr = serverAndPort[0], serverAndPort[1]
	} else {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing server host or port",
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

	params, err := url.ParseQuery(paramStr)
	if err != nil {
		return model.Outbound{}, &ParseError{
			Type:    ErrCannotParseParams,
			Raw:     proxy,
			Message: err.Error(),
		}
	}

	remarks, network, obfs, obfsPassword, pinSHA256, insecure, sni := params.Get("name"), params.Get("network"), params.Get("obfs"), params.Get("obfs-password"), params.Get("pinSHA256"), params.Get("insecure"), params.Get("sni")
	enableTLS := pinSHA256 != ""
	insecureBool := insecure == "1"

	result := model.Outbound{
		Type: "hysteria2",
		Tag:  remarks,
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
