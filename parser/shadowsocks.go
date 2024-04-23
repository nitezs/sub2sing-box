package parser

import (
	"net/url"
	"strings"
	"sub2sing-box/constant"
	"sub2sing-box/model"
	"sub2sing-box/util"
)

func ParseShadowsocks(proxy string) (model.Outbound, error) {
	if !strings.HasPrefix(proxy, constant.ShadowsocksPrefix) {
		return model.Outbound{}, &ParseError{Type: ErrInvalidPrefix, Raw: proxy}
	}

	proxy = strings.TrimPrefix(proxy, constant.ShadowsocksPrefix)
	urlParts := strings.SplitN(proxy, "@", 2)
	if len(urlParts) != 2 {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing character '@' in url",
			Raw:     proxy,
		}
	}

	var serverAndPort []string
	if !strings.Contains(urlParts[0], ":") {
		decoded, err := util.DecodeBase64(urlParts[0])
		if err != nil {
			return model.Outbound{}, &ParseError{
				Type:    ErrInvalidStruct,
				Message: "invalid base64 encoded",
				Raw:     proxy,
			}
		}
		urlParts[0] = decoded
	}
	credentials := strings.SplitN(urlParts[0], ":", 2)
	if len(credentials) != 2 {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing server host or port",
			Raw:     proxy,
		}
	}
	method, password := credentials[0], credentials[1]

	serverInfo := strings.SplitN(urlParts[1], "#", 2)
	serverAndPort = strings.SplitN(serverInfo[0], ":", 2)
	server, portStr := serverAndPort[0], serverAndPort[1]
	if len(serverInfo) != 2 {
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

	var remarks string
	if len(serverInfo) == 2 {
		unescape, err := url.QueryUnescape(serverInfo[1])
		if err != nil {
			return model.Outbound{}, &ParseError{
				Type:    ErrInvalidStruct,
				Message: "cannot unescape remarks",
				Raw:     proxy,
			}
		}
		remarks = strings.TrimSpace(unescape)
	} else {
		remarks = strings.TrimSpace(server + ":" + portStr)
	}

	result := model.Outbound{
		Type: "shadowsocks",
		Tag:  remarks,
		ShadowsocksOptions: model.ShadowsocksOutboundOptions{
			ServerOptions: model.ServerOptions{
				Server:     server,
				ServerPort: port,
			},
			Method:   method,
			Password: password,
		},
	}
	return result, nil
}
