package parser

import (
	"net/url"
	"strconv"
	"strings"
	"sub2sing-box/constant"
	"sub2sing-box/model"
)

func ParseHysteria(proxy string) (model.Outbound, error) {
	if !strings.HasPrefix(proxy, constant.HysteriaPrefix) {
		return model.Outbound{}, &ParseError{Type: ErrInvalidPrefix, Raw: proxy}
	}

	proxy = strings.TrimPrefix(proxy, constant.HysteriaPrefix)
	urlParts := strings.SplitN(proxy, "?", 2)
	if len(urlParts) != 2 {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing character '?' in url",
			Raw:     proxy,
		}
	}

	serverInfo := strings.SplitN(urlParts[0], ":", 2)
	if len(serverInfo) != 2 {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing server host or port",
			Raw:     proxy,
		}
	}
	server, portStr := serverInfo[0], serverInfo[1]

	port, err := ParsePort(portStr)
	if err != nil {
		return model.Outbound{}, err
	}

	params, err := url.ParseQuery(urlParts[1])
	if err != nil {
		return model.Outbound{}, &ParseError{
			Type:    ErrCannotParseParams,
			Raw:     proxy,
			Message: err.Error(),
		}
	}

	protocol, auth, insecure, upmbps, downmbps, obfs, alpnStr := params.Get("protocol"), params.Get("auth"), params.Get("insecure"), params.Get("upmbps"), params.Get("downmbps"), params.Get("obfs"), params.Get("alpn")
	insecureBool, err := strconv.ParseBool(insecure)
	if err != nil {
		insecureBool = false
	}

	var alpn []string
	alpnStr = strings.TrimSpace(alpnStr)
	if alpnStr != "" {
		alpn = strings.Split(alpnStr, ",")
	}

	remarks := server + ":" + portStr
	if params.Get("remarks") != "" {
		remarks = params.Get("remarks")
	}

	return model.Outbound{
		Type: "hysteria",
		Tag:  remarks,
		HysteriaOptions: model.HysteriaOutboundOptions{
			ServerOptions: model.ServerOptions{
				Server:     server,
				ServerPort: port,
			},
			Up:      upmbps,
			Down:    downmbps,
			Auth:    []byte(auth),
			Obfs:    obfs,
			Network: protocol,
			OutboundTLSOptionsContainer: model.OutboundTLSOptionsContainer{
				TLS: &model.OutboundTLSOptions{
					Enabled:  true,
					Insecure: insecureBool,
					ALPN:     alpn,
				},
			},
		},
	}, nil
}
