package parser

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
	"sub2sing-box/internal/model"
	"sub2sing-box/internal/util"
)

func ParseShadowsocks(proxy string) (model.Proxy, error) {
	if !strings.HasPrefix(proxy, "ss://") {
		return model.Proxy{}, errors.New("invalid ss Url")
	}
	parts := strings.SplitN(strings.TrimPrefix(proxy, "ss://"), "@", 2)
	if len(parts) != 2 {
		return model.Proxy{}, errors.New("invalid ss Url")
	}
	if !strings.Contains(parts[0], ":") {
		decoded, err := util.DecodeBase64(parts[0])
		if err != nil {
			return model.Proxy{}, errors.New("invalid ss Url" + err.Error())
		}
		parts[0] = decoded
	}
	credentials := strings.SplitN(parts[0], ":", 2)
	if len(credentials) != 2 {
		return model.Proxy{}, errors.New("invalid ss Url")
	}
	serverInfo := strings.SplitN(parts[1], "#", 2)
	serverAndPort := strings.SplitN(serverInfo[0], ":", 2)
	if len(serverAndPort) != 2 {
		return model.Proxy{}, errors.New("invalid ss Url")
	}
	port, err := strconv.Atoi(strings.TrimSpace(serverAndPort[1]))
	if err != nil {
		return model.Proxy{}, errors.New("invalid ss Url" + err.Error())
	}
	remarks := ""
	if len(serverInfo) == 2 {
		unescape, err := url.QueryUnescape(serverInfo[1])
		if err != nil {
			return model.Proxy{}, errors.New("invalid ss Url" + err.Error())
		}
		remarks = strings.TrimSpace(unescape)
	} else {
		remarks = strings.TrimSpace(serverAndPort[0])
	}
	method := credentials[0]
	password := credentials[1]
	server := strings.TrimSpace(serverAndPort[0])
	result := model.Proxy{
		Type: "shadowsocks",
		Tag:  remarks,
		Shadowsocks: model.Shadowsocks{
			Method:     method,
			Password:   password,
			Server:     server,
			ServerPort: uint16(port),
		},
	}
	return result, nil
}
