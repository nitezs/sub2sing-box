package parser

import (
	"net/url"
	"strconv"
	"strings"
	"sub2sing-box/constant"
	"sub2sing-box/model"
	"sub2sing-box/util"
)

func ParseShadowsocks(proxy string) (model.Outbound, error) {
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

	if link.Scheme+"://" != constant.ShadowsocksPrefix {
		return model.Outbound{}, &ParseError{Type: ErrInvalidPrefix, Raw: proxy}
	}

	port, err := strconv.Atoi(link.Port())
	if err != nil {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing server port",
			Raw:     proxy,
		}
	}

	user, err := util.DecodeBase64(link.User.Username())
	if user == "" {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing method and password",
			Raw:     proxy,
		}
	}
	methodAndPass := strings.SplitN(user, ":", 2)
	if len(methodAndPass) != 2 {
		return model.Outbound{}, &ParseError{
			Type:    ErrInvalidStruct,
			Message: "missing method and password",
			Raw:     proxy,
		}
	}

	query := link.Query()
	pluginStr := query.Get("plugin")
	plugin := ""
	options := ""
	if pluginStr != "" {
		arr := strings.SplitN(pluginStr, ";", 2)
		if len(arr) == 2 {
			plugin = arr[0]
			options = arr[1]
		}
	}

	result := model.Outbound{
		Type: "shadowsocks",
		Tag:  link.Fragment,
		ShadowsocksOptions: model.ShadowsocksOutboundOptions{
			ServerOptions: model.ServerOptions{
				Server:     server,
				ServerPort: uint16(port),
			},
			Method:        methodAndPass[0],
			Password:      methodAndPass[1],
			Plugin:        plugin,
			PluginOptions: options,
		},
	}
	return result, nil
}
