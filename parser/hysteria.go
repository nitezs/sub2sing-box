package parser

import (
	"fmt"
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

	query := link.Query()

	protocol, auth, insecure, upmbps, downmbps, obfs, alpnStr := query.Get("protocol"), query.Get("auth"), query.Get("insecure"), query.Get("upmbps"), query.Get("downmbps"), query.Get("obfs"), query.Get("alpn")
	insecureBool, err := strconv.ParseBool(insecure)
	if err != nil {
		insecureBool = false
	}

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
