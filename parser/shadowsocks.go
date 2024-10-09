package parser

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/nitezs/sub2sing-box/constant"
	"github.com/nitezs/sub2sing-box/model"
	"github.com/nitezs/sub2sing-box/util"
	"github.com/sagernet/sing-box/option"
)

func ParseShadowsocks(proxy string) (model.Outbound, error) {
	if !strings.HasPrefix(proxy, constant.ShadowsocksPrefix) {
		return model.Outbound{}, &ParseError{Type: ErrInvalidPrefix, Raw: proxy}
	}
	needDecode := true
	if !strings.Contains(proxy, "@") {
		s := strings.SplitN(proxy, "#", 2)
		d, err := util.DecodeBase64(strings.TrimPrefix(s[0], "ss://"))
		if err != nil {
			return model.Outbound{}, &ParseError{
				Type:    ErrInvalidStruct,
				Message: "url parse error",
				Raw:     proxy,
			}
		}
		if len(s) == 2 {
			proxy = "ss://" + d + "#" + s[1]
		} else {
			proxy = "ss://" + d
		}
		needDecode = false
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
			Type: ErrInvalidStruct,
			Raw:  proxy,
		}
	}

	method := ""
	password := ""
	if needDecode {
		user, err := util.DecodeBase64(link.User.Username())
		if err != nil {
			return model.Outbound{}, &ParseError{
				Type:    ErrInvalidStruct,
				Message: "missing method and password",
				Raw:     proxy,
			}
		}
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
		method = methodAndPass[0]
		password = methodAndPass[1]
	} else {
		method = link.User.Username()
		password, _ = link.User.Password()
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

	remarks := link.Fragment
	if remarks == "" {
		remarks = fmt.Sprintf("%s:%s", server, portStr)
	}
	remarks = strings.TrimSpace(remarks)

	result := model.Outbound{
		Outbound: option.Outbound{
			Type: "shadowsocks",
			Tag:  remarks,
			ShadowsocksOptions: option.ShadowsocksOutboundOptions{
				ServerOptions: option.ServerOptions{
					Server:     server,
					ServerPort: port,
				},
				Method:        method,
				Password:      password,
				Plugin:        plugin,
				PluginOptions: options,
			},
		},
	}
	return result, nil
}
